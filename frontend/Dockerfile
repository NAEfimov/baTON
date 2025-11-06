#DEV SETUP - UNCOMMENT THE VOLUMES IN DOCKER-COMPOSE IN CASE OF IT NOT WORKING

# ----------------------------
# Stage 1: Dev Environment
FROM node:22 AS dev
WORKDIR /app

# Copy dependencies
COPY package*.json ./
RUN npm ci

# Copy the rest of the app
COPY . .

# Environment for Vite
ENV HOST=0.0.0.0
ENV PORT=8081

EXPOSE 8081

# Run Vite development server (hot reload)
CMD ["npm", "run", "dev"]

# Stage 1: Build
# FROM node:22 AS builder
# WORKDIR /app

# COPY package*.json ./
# RUN npm ci

# COPY . .
# RUN npm run build

# # ----------------------------
# # Stage 2: Production
# FROM node:22-alpine
# WORKDIR /app

# # Copy only what's needed
# COPY --from=builder /app/dist ./dist
# COPY --from=builder /app/package*.json ./

# # Install only production dependencies
# RUN npm ci --omit=dev

# ENV HOST=0.0.0.0
# EXPOSE 8081

# CMD ["npx", "serve", "-s", "dist", "--listen", "8081", "--no-compression"]