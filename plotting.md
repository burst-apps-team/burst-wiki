The process in Burst that prepares your hard disk space by generating and storing shabal256 hashes for [mining](mining.md) is called plotting. The software used for this is called [plotter](plotter.md). Plotter software computes hashes using the cryptographic hash function (Shabal-256) and stores them in plot files on hard disks. Mining software retrieves these hashes to find values that can be used to forge blocks.

Creating plot files

Plot files are bound to the Burst account ID so it is impossible for users with different account IDs to generate identical plot files. Plot files can be created (with the same account ID) on other machines and transferred back to your mining computer. It is a common strategy to plot hard drives using computers with faster CPUs or GPUs and then mine the disks using computers with slower CPUs.

When creating plot files, it is important to avoid overlapping plots. Overlapping plots (duplicate plots) will not cause a malfunction, but will waste space that would otherwise be productive if containing unique plots. Duplicated plots contribute no added value.

If you will be using plotting software that uses a GPU, it is recommended that the GPU be used exclusively for plotting. Plotting is by nature a temporary process, so the GPU can be placed back into service for gaming or mining other coins after the process is complete. Gaming can occur during the plotting process, however, performance may suffer. Mining other coins during the plotting process is likely to result in corrupted plots. This is a general rule as some plotting software may have innovations that prevent this.

An easy to use plotter, XPlotter, is included in the windows based QBundle. QBundle also included a dynamic plotter that will manage your computers disk space so that available space is dedicated to mining but a dynamic amount of free space is maintained for normal usage.

TurboPlotter9000 is a proven plotter with increased speed, GPU support, and an easy to use interface that helps to prevent overlap. For fastest plotting speed, it can plot to an SMR drive and then transfer the created files to a hard disk seamlessly.

Other plotters have been developed for various operating systems and use cases.

Details for selecting plotting software can be found here:

[Available Plotting Software](burst-software-plotting-software.md)
