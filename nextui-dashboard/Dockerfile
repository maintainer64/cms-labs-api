FROM docker.io/library/nginx:1.31.6-alpine3.24

COPY dist /usr/share/nginx/html/
COPY nginx /etc/nginx/
COPY 20-envsubst-on-templates.sh /docker-entrypoint.d

# Открываем порт 80
EXPOSE 80
EXPOSE 443
