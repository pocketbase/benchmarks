# https://fly.io/docs/rails/cookbooks/deploy/#enabling-swap
fallocate -l 512M /swapfile
chmod 0600 /swapfile
mkswap /swapfile
echo 10 > /proc/sys/vm/swappiness
swapon /swapfile
echo 1 > /proc/sys/vm/overcommit_memory

# run the application
/pb/pocketbase serve --http=0.0.0.0:8090
