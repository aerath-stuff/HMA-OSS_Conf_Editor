package consts

/**
 * Defines the GID for the group that allows write access to the internal media storage.
 */
const SDCARD_RW_GID = 1015

/**
 * Defines the GID for the group that allows write access to the internal media storage.
 */
const MEDIA_RW_GID = 1023

/**
 * Access to installed package details
 */
const PACKAGE_INFO_GID = 1032

/*
 * GID that gives access to USB OTG (unreliable) volumes on /mnt/media_rw/<vol name>
 */
const EXTERNAL_STORAGE_GID = 1077

/**
 * GID that gives write access to app-private data directories on external
 * storage (used on devices without sdcardfs only).
 */
const EXT_DATA_RW_GID = 1078

/**
 * GID that gives write access to app-private OBB directories on external
 * storage (used on devices without sdcardfs only).
 */
const EXT_OBB_RW_GID = 1079

/**
 * GID that corresponds to the INTERNET permission.
 * Must match the value of AID_INET.
 */
const INET_GID = 3003

/**
 * Defines the gid shared by all applications running under the same profile.
 */
const SHARED_USER_GID = 9997

var ZygotePermissions = map[int]string{
	SDCARD_RW_GID:        "SDCARD_RW_GID",
	MEDIA_RW_GID:         "MEDIA_RW_GID",
	PACKAGE_INFO_GID:     "PACKAGE_INFO_GID",
	EXTERNAL_STORAGE_GID: "EXTERNAL_STORAGE_GID",
	EXT_DATA_RW_GID:      "EXT_DATA_RW_GID",
	EXT_OBB_RW_GID:       "EXT_OBB_RW_GID",
	INET_GID:             "INET_GID",
	SHARED_USER_GID:      "SHARED_USER_GID",
}
