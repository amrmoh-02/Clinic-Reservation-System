# Use an official MongoDB image as a parent image
FROM mongo:7.0

# Expose port 27017 to the outside world
EXPOSE 27017

# Command to run MongoDB
CMD ["mongod"]