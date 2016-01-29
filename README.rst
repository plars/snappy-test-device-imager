Snappy Test Device Imager
#########################

Snappy Test Device Imager is a tool for booting a minimal initramfs on a system
that can be used for installing snappy to a local disk.  It consists of two
components: A service written in go that accepts connections on port 8989, and
the scripts use to build an initramfs that can be booted, and automatically
starts the service on the remote machine.

Building the Imager Tool
========================

First, you'll need to install golang-go on your system, if it's not installed
already.

To build the go binary, simply cd to the snappy-test-device-imager directory
under this project, and type::

	$ go build

If everything worked properly, you should be left with a binary named
'snappy-test-device-imager' in that directory.

Building the initramfs
======================

To build the initramfs, you'll need to get the following dependencies:
initramfs-tools, netcat-traditional

.. note:: Make sure to use netcat-traditional rather than netcat-openbsd

Once you have installed the dependencies, run the following from the project
root directory::

	$ mkinitramfs -d initramfs -o init.img

This will build the initramfs, and put it in init.img in your local directory.

To boot this on your system, you'll need to set up pxe booting in your
environment, and copy the kernel from your system (unless you specified a
different one) and the init.img to the tftp server, and specify those in the pxe
config for the system you are trying to boot.

Using the Snappy Test Device Imager Tool
========================================

You will first need to build a snappy image, gzip it, and place it on some
system that can access by the target device over the network.  You can use
a tool such as netcat to stream the image over the network using something
like::

	$ cat /tmp/snappy.img.gz |nc -l 9999

In the above example, netcat will listen for connections on port 9999 and
stream the image to a client that connects.

Next, you'll need the ip address of the system that has booted the initramfs
produced by snappy-test-device-imager. You will also need to know which
block device on the target system should be used for installation.

To start the installation, you can use curl or any similar tool to make a
request like this::

	$ curl TARGET_IP:8989/writeimage?server=IMAGE_HOST_IP:PORT\&dev=/dev/sda

In the example above, we used /dev/sda for the target block device. Make sure
to replace this with an appropriate block device for your target.

Additional Utilities
====================

A small command line tool under utils/pxe-setup is provided for conveniently
creating and destroying pxe config files for tftp/netboot.  It should be
installed in the default path on your pxeboot server, and is used by
the snappy device agent for netboot.
