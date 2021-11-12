FROM postgres:latest

# Localeを日本語にセット
RUN export LC_ALL="ja_JP.UTF-8"
RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8

# 言語を日本語にセット
ENV LANG ja_JP.UTF-8
