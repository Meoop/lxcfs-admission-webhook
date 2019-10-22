package lxcfs

import (
	corev1 "k8s.io/api/core/v1"
)

// volumeMounts:
//   - mountPath: /proc/cpuinfo
//     name: lxcfs-cpuinfo
//   - mountPath: /sys/devices/system/cpu/online
//     name: lxcfs-onlinecpu
//   - mountPath: /proc/meminfo
//     name: lxcfs-meminfo
//   - mountPath: /proc/diskstats
//     name: lxcfs-diskstats
//   - mountPath: /proc/stat
//     name: lxcfs-stat
//   - mountPath: /proc/swaps
//     name: lxcfs-swaps
//   - mountPath: /proc/uptime
//	   name: lxcfs-uptime
var volumeMountsTemplate = []corev1.VolumeMount{

	{
		Name:      "lxcfs-cpuinfo",
		MountPath: "/proc/cpuinfo",
	},
	{
		Name:      "lxcfs-onlinecpu",
		MountPath: "/sys/devices/system/cpu/online",
	},
	{
		Name:      "lxcfs-meminfo",
		MountPath: "/proc/meminfo",
	},
	{
		Name:      "lxcfs-diskstats",
		MountPath: "/proc/diskstats",
	},
	{
		Name:      "lxcfs-stat",
		MountPath: "/proc/stat",
	},
	{
		Name:      "lxcfs-swaps",
		MountPath: "/proc/swaps",
	},
	{
		Name:      "lxcfs-uptime",
		MountPath: "/proc/uptime",
	},
}

// volumes:
//   - name: lxcfs-cpuinfo
//     hostPath:
//       path: /var/lib/lxcfs/proc/cpuinfo
//   - name: lxcfs-onlinecpu
//     hostPath:
//       path: /var/lib/lxcfs/sys/devices/system/cpu/online
//   - name: lxcfs-meminfo
//     hostPath:
//      path: /var/lib/lxcfs/proc/meminfo
//   - name: lxcfs-diskstats
//     hostPath:
//       path: /var/lib/lxcfs/proc/diskstats
//   - name: lxcfs-stat
//     hostPath:
//       path: /var/lib/lxcfs/proc/stat
//   - name: lxcfs-swaps
//     hostPath:
//       path: /var/lib/lxcfs/proc/swaps
//   - name: lxcfs-uptime
//     hostPath:
//       path: /var/lib/lxcfs/proc/uptime
var volumesTemplate = []corev1.Volume{
	{
		Name: "lxcfs-cpuinfo",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/proc/cpuinfo",
			},
		},
	},
	{
		Name: "lxcfs-onlinecpu",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/sys/devices/system/cpu/online",
			},
		},
	},
	{
		Name: "lxcfs-meminfo",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/proc/meminfo",
			},
		},
	},
	{
		Name: "lxcfs-diskstats",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/proc/diskstats",
			},
		},
	},
	{
		Name: "lxcfs-stat",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/proc/stat",
			},
		},
	},
	{
		Name: "lxcfs-swaps",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/proc/swaps",
			},
		},
	},
	{
		Name: "lxcfs-uptime",
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/var/lib/lxcfs/proc/uptime",
			},
		},
	},
}
