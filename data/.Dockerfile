# .Dockerfile
FROM neo4j:latest

# Set environment variables for Neo4j authentication
ENV NEO4J_AUTH=${NEO4J_AUTH}

# Expose ports for Neo4j Browser and Bolt protocol
EXPOSE 7474 7687

# Start Neo4j
CMD ["neo4j", "console"]