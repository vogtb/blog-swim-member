# -*- mode: ruby -*-
# vi: set ft=ruby :

IP_ADDRESS=ENV['IP']
BOX_NAME=ENV['NAME']

Vagrant.configure("2") do |config|
  # setup box
  config.vm.box = BOX_NAME
  config.vm.box_url = "http://files.vagrantup.com/precise64.box"

  # name it
  config.vm.hostname = BOX_NAME

  # setup VirtualBox as the provider
  config.vm.provider :virtualbox do |vb|
    vb.name = BOX_NAME
    vb.customize ["modifyvm", :id, "--memory", "512"]
    vb.customize ["modifyvm", :id, "--cpus", "2"]
    vb.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
    vb.customize ["modifyvm", :id, "--natdnsproxy1", "on"]
  end

  config.vm.box_check_update = false

  # sync all folders inside project src directory
  config.vm.synced_folder "../../", "/mnt/src", mount_options: ["dmode=777", "fmode=666"]

  # give access to public network, specify bridge
  config.vm.network "private_network", ip: IP_ADDRESS

end
