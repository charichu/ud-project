Modify hosts file for caddy config see caddyfile
127.0.0.1 backend

linux setup: create a user with sudo rights

ufw setup:
 ufw allow ssh
 ufw allow http
 ufw allow https
 ufw allow 2377/tcp
 ufw allow 7946/tcp
 ufw allow 7946/udp
 ufw allow 4789/udp
 ufw allow 8025/tcp
 ufw enable