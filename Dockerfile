FROM archlinux:base
ENV TERM="xterm-256"
RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
	&& echo "Server = https://mirrors.tuna.tsinghua.edu.cn/archlinux/\$repo/os/\$arch" > /etc/pacman.d/mirrorlist \
	&& pacman -Syy \
	&& pacman -S --needed --noconfirm \
	zsh which vim curl sed awk openssh sudo git gcc make \
	&& echo "PermitRootLogin yes" >> /etc/ssh/sshd_config \
	&& mkdir /root/.ssh \
	&& mkdir /root/Workspace
EXPOSE 22
WORKDIR /root
CMD [ ! -f /etc/ssh/ssh_host_rsa_key ] && ssh-keygen -A; /bin/sshd -D