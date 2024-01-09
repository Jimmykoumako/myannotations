# Use the official Node.js image as a base image
FROM node:latest

# Set the working directory
WORKDIR /usr/src/app

# Copy package.json and package-lock.json to the working directory
COPY package*.json ./

# Install Angular CLI globally
RUN npm install -g @angular/cli

# Install ng globally
RUN npm install -g ng

# Install dependencies
RUN npm install

# Copy the application files to the working directory
COPY . .

# Expose the default Angular port
EXPOSE 4200

# Start the Angular application
CMD ["npm", "start"]
