FROM node:latest


WORKDIR /app

# Install dependencies
COPY package.json ./
COPY package-lock.json ./
RUN npm install

# Copy the source code
COPY . ./

# Expose the development server port
EXPOSE 3000

# Start the development server
CMD ["npm", "start"]