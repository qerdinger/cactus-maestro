docker build -t cmaestro_db .

docker run -d \
  --name redis \
  --restart unless-stopped \
  -p 6379:6379 \
  -v redis-data:/var/lib/redis \
  cmaestro_db