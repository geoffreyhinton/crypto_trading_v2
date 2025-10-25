# Bitcoin RPC Explorer Dockerfile
FROM node:18-alpine

# Set working directory
WORKDIR /app

# Install git and other dependencies
RUN apk add --no-cache git

# Clone the btc-rpc-explorer repository
RUN git clone https://github.com/janoside/btc-rpc-explorer.git .

# Install dependencies
RUN npm install

# Expose port
EXPOSE 3002

# Set environment variables
ENV BTCEXP_HOST=0.0.0.0
ENV BTCEXP_PORT=3002
ENV NODE_ENV=production

# Start the application
CMD ["npm", "start"]