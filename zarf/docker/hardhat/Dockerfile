# Use an official Node.js runtime as a parent image
FROM node:18

# Set the working directory in the container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Install Hardhat globally
RUN npm install -g hardhat

# Expose the port Hardhat uses
EXPOSE 8545

# Run Hardhat node
CMD ["npx", "hardhat", "node"]
