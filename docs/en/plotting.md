Introduction to Burst Plotting
------------------------------

The process that prepares your hard drive space for [mining](mining.md) by generating and storing shabal256 hashes is called plotting. The software used for this is called a plotter. Plotter software computes hashes using the cryptographic hash function (Shabal-256) and stores them in plot files on hard disks. Mining software retrieves these hashes to find values that can be used to forge blocks.

Plot files are bound to the Burst account ID so it is impossible for users with different account IDs to generate identical plot files. Plot files can be created (with the same account ID) on other machines and transferred back to your mining computer. It is a common strategy to plot hard drives using computers with faster CPUs or GPUs and then mine the disks using computers with slower CPUs.

When creating plot files, it is important to avoid overlapping plots. Overlapping plots (duplicate plots) will not cause a malfunction, but will waste space that would otherwise be productive if containing unique plots. Duplicated plots contribute no added value.

If you will be using plotting software that uses a GPU, it is recommended that the GPU be used exclusively for plotting. Plotting is by nature a temporary process, so the GPU can be placed back into service for gaming or mining other coins after the process is complete. Gaming can occur during the plotting process, however, performance may suffer. Mining other coins using the same GPU during the plotting process is likely to result in corrupted plots. This is a general rule as some plotting software may have innovations that prevent this.

Plotting Software
-----------------

An easy to use plotter, XPlotter, is included in the windows based QBundle. QBundle also includes a dynamic plotter that will manage your computers disk space so that available space is dedicated to mining but a dynamic amount of free space is maintained for normal usage.

Other plotters have been developed for various operating systems and use cases.

Details for selecting plotting software can be found [here](burst-software-plotting-software.md). All documentation regarding the use of plotters is available on the software's website and/or on Github.
