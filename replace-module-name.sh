#!/bin/bash
set -e

# Get the old module name from the previous commit
git show HEAD:go.mod | grep '^module ' | awk '{print $2}' > /tmp/old_module_name.txt
OLD_MODULE=$(cat /tmp/old_module_name.txt)

# Get the GitHub remote URL and extract the module name
REMOTE_URL=$(git remote get-url origin)
if [[ "$REMOTE_URL" =~ github.com[:/](.+)\.git ]]; then
  NEW_MODULE="github.com/${BASH_REMATCH[1]}"
else
  echo "Could not extract module name from remote URL: $REMOTE_URL"
  exit 1
fi

# Update go.mod with the new module name
sed -i '' "s|^module .*|module $NEW_MODULE|" go.mod

if [ "$OLD_MODULE" = "$NEW_MODULE" ]; then
  echo "Module name has not changed."
  exit 0
fi

# App name is the last segment of the new module name
APP_NAME=$(basename "$NEW_MODULE")
# Replace hyphens with underscores for DB, user, and password
APP_DB_NAME="${APP_NAME//-/_}"

# Update Go imports
git grep -lz "$OLD_MODULE" -- '*.go' | xargs -0 sed -i '' "s|$OLD_MODULE|$NEW_MODULE|g"

# Update .golangci.yaml for goimports and gci sections
sed -i '' "s|$OLD_MODULE|$NEW_MODULE|g" .golangci.yaml

# Update postgres params in docker-compose.yml
sed -i '' "s|POSTGRES_USER: .*|POSTGRES_USER: ${APP_DB_NAME}_user|g" docker-compose.yml
sed -i '' "s|POSTGRES_PASSWORD: .*|POSTGRES_PASSWORD: ${APP_DB_NAME}_password|g" docker-compose.yml
sed -i '' "s|POSTGRES_DB: .*|POSTGRES_DB: ${APP_DB_NAME}|g" docker-compose.yml

# Update postgresConn in config.yaml
sed -i '' "s|postgresql://.*:.*@localhost:5432/.*|postgresql://${APP_DB_NAME}_user:${APP_DB_NAME}_password@localhost:5432/${APP_DB_NAME}|g" config.yaml

go mod tidy

echo "Module name set to $NEW_MODULE and all related files updated."
