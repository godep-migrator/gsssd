# -*- mode: ruby -*-
# vi: set ft=ruby :

$build = <<END
sudo apt-get update -y
sudo apt-get install -y build-essential git

echo "Host git.internal.justin.tv" >> ~/.ssh/config
echo "  StrictHostKeyChecking no" >> ~/.ssh/config

git clone git@git.internal.justin.tv:release/debpkg /tmp/debpkg
sudo dpkg -i /tmp/debpkg/twitch-golang_1.2-1_amd64.deb

mkdir ~/build
cp -R ~/src/* ~/build
cd ~/build
export GOPATH=~/go
make clean prep
sudo chown -R root.root deb/
sudo make deb
cp ./gsssd_*.deb ~/src
END

# Vagrantfile API/syntax version. Don't touch unless you know what you're doing!
VAGRANTFILE_API_VERSION = "2"

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "precise64"
  config.vm.box_url = "http://files.vagrantup.com/precise64.box"
  config.ssh.forward_agent = true
  config.vm.provider :virtualbox do |vb|
    vb.customize ["modifyvm", :id, "--memory", "2048"]
  end
  config.vm.synced_folder ENV['GOPATH'], "/home/vagrant/go"
  config.vm.synced_folder ".", "/home/vagrant/src"
  config.vm.provision :shell, inline: $build, privileged: false
end
