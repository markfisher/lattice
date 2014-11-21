VAGRANTFILE_API_VERSION = "2"

if ENV["VAGRANT_DIEGO_EDGE_TAR_PATH"].nil?
    puts "VAGRANT_DIEGO_EDGE_TAR_PATH env var must be set to where the diego-edge tar lives when inside the vagrant box eg /vagrant/diego-edge-latest.tgz"
    exit 1
end

Vagrant.configure(VAGRANTFILE_API_VERSION) do |config|
  config.vm.box = "ubuntu/trusty64"
  config.vm.network "private_network", ip: "192.168.11.11"
  config.vm.provision "shell" do |s|
    s.path = "install_from_tar"
    s.args = ENV["VAGRANT_DIEGO_EDGE_TAR_PATH"]
  end

  config.vm.provider "virtualbox" do |v|
    # dns resolution appears to be very slow in some environments; this fixes it
    v.customize ["modifyvm", :id, "--natdnshostresolver1", "on"]
  end
end