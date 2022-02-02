docker build -t leoryu/archlinux:dev .
docker run -d -v $HOME/Workspace/test:/root/Workspace -p 50000:22 --name dev leoryu/archlinux:dev
docker exec -ti  dev bash
