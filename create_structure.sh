#!/bin/bash

# Define your project name
PROJECT_NAME="myannotations"

# Create necessary directories
mkdir -p $PROJECT_NAME/backend/controllers
mkdir -p $PROJECT_NAME/backend/models
mkdir -p $PROJECT_NAME/backend/config
mkdir -p $PROJECT_NAME/backend/routers
mkdir -p $PROJECT_NAME/frontend/src/app/components
mkdir -p $PROJECT_NAME/frontend/src/app/services
mkdir -p $PROJECT_NAME/migrations
mkdir -p $PROJECT_NAME/utils
mkdir -p $PROJECT_NAME/.github/workflows

# Create necessary files
touch $PROJECT_NAME/main.go
touch $PROJECT_NAME/go.mod
touch $PROJECT_NAME/go.sum
touch $PROJECT_NAME/backend/controllers/user_controller.go
touch $PROJECT_NAME/backend/controllers/book_controller.go
touch $PROJECT_NAME/backend/controllers/note_controller.go
touch $PROJECT_NAME/backend/controllers/highlight_controller.go
touch $PROJECT_NAME/backend/controllers/connection_controller.go
touch $PROJECT_NAME/backend/models/models.go
touch $PROJECT_NAME/backend/models/user.go
touch $PROJECT_NAME/backend/models/book.go
touch $PROJECT_NAME/backend/models/text.go
touch $PROJECT_NAME/backend/models/annotation.go
touch $PROJECT_NAME/backend/models/connection.go
touch $PROJECT_NAME/backend/config/database.go
touch $PROJECT_NAME/backend/routers/router.go
touch $PROJECT_NAME/frontend/src/app/components/book-list.component.ts
touch $PROJECT_NAME/frontend/src/app/components/annotation-view.component.ts
touch $PROJECT_NAME/frontend/src/app/services/data.service.ts
touch $PROJECT_NAME/migrations/20240101000001_create_tables.up.sql
touch $PROJECT_NAME/migrations/20240101000001_create_tables.down.sql
touch $PROJECT_NAME/utils/helpers.go
touch $PROJECT_NAME/.dockerignore
touch $PROJECT_NAME/Dockerfile
touch $PROJECT_NAME/.gitignore
touch $PROJECT_NAME/.github/workflows/ci.yml

# Write content to files
echo "module $PROJECT_NAME" > $PROJECT_NAME/go.mod
echo "/* Your Go code here */" > $PROJECT_NAME/main.go
echo "/* Your Dockerfile content here */" > $PROJECT_NAME/Dockerfile
echo "/node_modules" > $PROJECT_NAME/.dockerignore
echo "/* Your .gitignore content here */" > $PROJECT_NAME/.gitignore
echo "/* Your GitHub Actions workflow content here */" > $PROJECT_NAME/.github/workflows/ci.yml

# Create frontend files
touch $PROJECT_NAME/frontend/angular.json
touch $PROJECT_NAME/frontend/package.json
touch $PROJECT_NAME/frontend/tsconfig.json
echo "/* Your Angular configuration content here */" > $PROJECT_NAME/frontend/angular.json
echo "/* Your package.json content here */" > $PROJECT_NAME/frontend/package.json
echo "/* Your tsconfig.json content here */" > $PROJECT_NAME/frontend/tsconfig.json

# Initialize Git repository
git init
git add .
git commit -m "Initial commit"

# Create a GitHub repository (replace "your_username" and "your_access_token" with your GitHub username and token)
curl -u "Jimmykoumako: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" https://api.github.com/Jimmykoumako/repos -d '{"name":"'$PROJECT_NAME'"}'

# Set the remote URL and push to GitHub
git remote add origin https://github.com/Jimmykoumako/$PROJECT_NAME.git
git branch -M main
git push -u origin main

# Inform the user
echo "Project initialized and pushed to GitHub successfully!"
