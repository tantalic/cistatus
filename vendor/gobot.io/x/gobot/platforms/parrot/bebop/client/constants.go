package client

const (
	// libARNetworkAL/Includes/libARNetworkAL/ARNETWORKAL_Manager.h
	ARNETWORKAL_MANAGER_DEFAULT_ID_MAX uint16 = 256

	// ARNETWORKAL_Frame_t identifiers
	BD_NET_CD_NONACK_ID     byte = 10
	BD_NET_CD_ACK_ID        byte = 11
	BD_NET_CD_EMERGENCY_ID  byte = 12
	BD_NET_CD_VIDEO_ACK_ID  byte = 13
	BD_NET_DC_VIDEO_DATA_ID byte = 125
	BD_NET_DC_EVENT_ID      byte = 126
	BD_NET_DC_NAVDATA_ID    byte = 127

	// eARCOMMANDS_ID_PROJECT
	ARCOMMANDS_ID_PROJECT_COMMON   byte = 0
	ARCOMMANDS_ID_PROJECT_ARDRONE3 byte = 1

	// eARCOMMANDS_ID_ARDRONE3_CLASS
	ARCOMMANDS_ID_ARDRONE3_CLASS_PILOTING              byte = 0
	ARCOMMANDS_ID_ARDRONE3_CLASS_ANIMATIONS            byte = 5
	ARCOMMANDS_ID_ARDRONE3_CLASS_CAMERA                byte = 1
	ARCOMMANDS_ID_ARDRONE3_CLASS_MEDIARECORD           byte = 7
	ARCOMMANDS_ID_ARDRONE3_CLASS_MEDIARECORDSTATE      byte = 8
	ARCOMMANDS_ID_ARDRONE3_CLASS_MEDIARECORDEVENT      byte = 3
	ARCOMMANDS_ID_ARDRONE3_CLASS_PILOTINGSTATE         byte = 4
	ARCOMMANDS_ID_ARDRONE3_CLASS_NETWORK               byte = 13
	ARCOMMANDS_ID_ARDRONE3_CLASS_NETWORKSTATE          byte = 14
	ARCOMMANDS_ID_ARDRONE3_CLASS_PILOTINGSETTINGS      byte = 2
	ARCOMMANDS_ID_ARDRONE3_CLASS_PILOTINGSETTINGSSTATE byte = 6
	ARCOMMANDS_ID_ARDRONE3_CLASS_SPEEDSETTINGS         byte = 11
	ARCOMMANDS_ID_ARDRONE3_CLASS_SPEEDSETTINGSSTATE    byte = 12
	ARCOMMANDS_ID_ARDRONE3_CLASS_NETWORKSETTINGS       byte = 9
	ARCOMMANDS_ID_ARDRONE3_CLASS_NETWORKSETTINGSSTATE  byte = 10
	ARCOMMANDS_ID_ARDRONE3_CLASS_SETTINGS              byte = 15
	ARCOMMANDS_ID_ARDRONE3_CLASS_SETTINGSSTATE         byte = 16
	ARCOMMANDS_ID_ARDRONE3_CLASS_DIRECTORMODE          byte = 17
	ARCOMMANDS_ID_ARDRONE3_CLASS_DIRECTORMODESTATE     byte = 18
	ARCOMMANDS_ID_ARDRONE3_CLASS_PICTURESETTINGS       byte = 19
	ARCOMMANDS_ID_ARDRONE3_CLASS_PICTURESETTINGSSTATE  byte = 20
	ARCOMMANDS_ID_ARDRONE3_CLASS_MEDIASTREAMING        byte = 21
	ARCOMMANDS_ID_ARDRONE3_CLASS_MEDIASTREAMINGSTATE   byte = 22
	ARCOMMANDS_ID_ARDRONE3_CLASS_GPSSETTINGS           byte = 23
	ARCOMMANDS_ID_ARDRONE3_CLASS_GPSSETTINGSSTATE      byte = 24
	ARCOMMANDS_ID_ARDRONE3_CLASS_CAMERASTATE           byte = 25
	ARCOMMANDS_ID_ARDRONE3_CLASS_ANTIFLICKERING        byte = 29
	ARCOMMANDS_ID_ARDRONE3_CLASS_ANTIFLICKERINGSTATE   byte = 30

	// eARCOMMANDS_ID_ARDRONE3_PILOTINGSTATE_CMD
	ARCOMMANDS_ID_ARDRONE3_PILOTINGSTATE_CMD_FLATTRIMCHANGED          byte = 0
	ARCOMMANDS_ID_ARDRONE3_PILOTINGSTATE_CMD_FLYINGSTATECHANGED       byte = 1
	ARCOMMANDS_ID_ARDRONE3_PILOTINGSTATE_CMD_ALERTSTATECHANGED        byte = 2
	ARCOMMANDS_ID_ARDRONE3_PILOTINGSTATE_CMD_NAVIGATEHOMESTATECHANGED byte = 3
	ARCOMMANDS_ID_ARDRONE3_PILOTINGSTATE_CMD_POSITIONCHANGED          byte = 4
	ARCOMMANDS_ID_ARDRONE3_PILOTINGSTATE_CMD_SPEEDCHANGED             byte = 5
	ARCOMMANDS_ID_ARDRONE3_PILOTINGSTATE_CMD_ATTITUDECHANGED          byte = 6
	ARCOMMANDS_ID_ARDRONE3_PILOTINGSTATE_CMD_AUTOTAKEOFFMODECHANGED   byte = 7
	ARCOMMANDS_ID_ARDRONE3_PILOTINGSTATE_CMD_ALTITUDECHANGED          byte = 8
	ARCOMMANDS_ID_ARDRONE3_PILOTINGSTATE_CMD_MAX                      byte = 9

	// eARCOMMANDS_ID_ARDRONE3_ANIMATIONS_CMD;
	ARCOMMANDS_ID_ARDRONE3_ANIMATIONS_CMD_FLIP byte = 0
	ARCOMMANDS_ID_ARDRONE3_ANIMATIONS_CMD_MAX  byte = 1

	// eARCOMMANDS_ARDRONE3_PILOTINGSTATE_FLYINGSTATECHANGED_STATE;
	ARCOMMANDS_ARDRONE3_PILOTINGSTATE_FLYINGSTATECHANGED_STATE_LANDED    byte = 0
	ARCOMMANDS_ARDRONE3_PILOTINGSTATE_FLYINGSTATECHANGED_STATE_TAKINGOFF byte = 1
	ARCOMMANDS_ARDRONE3_PILOTINGSTATE_FLYINGSTATECHANGED_STATE_HOVERING  byte = 2
	ARCOMMANDS_ARDRONE3_PILOTINGSTATE_FLYINGSTATECHANGED_STATE_FLYING    byte = 3
	ARCOMMANDS_ARDRONE3_PILOTINGSTATE_FLYINGSTATECHANGED_STATE_LANDING   byte = 4
	ARCOMMANDS_ARDRONE3_PILOTINGSTATE_FLYINGSTATECHANGED_STATE_EMERGENCY byte = 5
	ARCOMMANDS_ARDRONE3_PILOTINGSTATE_FLYINGSTATECHANGED_STATE_MAX       byte = 6

	// eARCOMMANDS_ARDRONE3_ANIMATIONS_FLIP_DIRECTION;
	ARCOMMANDS_ARDRONE3_ANIMATIONS_FLIP_DIRECTION_FRONT byte = 0
	ARCOMMANDS_ARDRONE3_ANIMATIONS_FLIP_DIRECTION_BACK  byte = 1
	ARCOMMANDS_ARDRONE3_ANIMATIONS_FLIP_DIRECTION_RIGHT byte = 2
	ARCOMMANDS_ARDRONE3_ANIMATIONS_FLIP_DIRECTION_LEFT  byte = 3
	ARCOMMANDS_ARDRONE3_ANIMATIONS_FLIP_DIRECTION_MAX   byte = 4

	// eARCOMMANDS_ID_COMMON_CLASS
	ARCOMMANDS_ID_COMMON_CLASS_NETWORK             byte = 0
	ARCOMMANDS_ID_COMMON_CLASS_NETWORKEVENT        byte = 1
	ARCOMMANDS_ID_COMMON_CLASS_SETTINGS            byte = 2
	ARCOMMANDS_ID_COMMON_CLASS_SETTINGSSTATE       byte = 3
	ARCOMMANDS_ID_COMMON_CLASS_COMMON              byte = 4
	ARCOMMANDS_ID_COMMON_CLASS_COMMONSTATE         byte = 5
	ARCOMMANDS_ID_COMMON_CLASS_OVERHEAT            byte = 6
	ARCOMMANDS_ID_COMMON_CLASS_OVERHEATSTATE       byte = 7
	ARCOMMANDS_ID_COMMON_CLASS_CONTROLLERSTATE     byte = 8
	ARCOMMANDS_ID_COMMON_CLASS_WIFISETTINGS        byte = 9
	ARCOMMANDS_ID_COMMON_CLASS_WIFISETTINGSSTATE   byte = 10
	ARCOMMANDS_ID_COMMON_CLASS_MAVLINK             byte = 11
	ARCOMMANDS_ID_COMMON_CLASS_MAVLINKSTATE        byte = 12
	ARCOMMANDS_ID_COMMON_CLASS_CALIBRATION         byte = 13
	ARCOMMANDS_ID_COMMON_CLASS_CALIBRATIONSTATE    byte = 14
	ARCOMMANDS_ID_COMMON_CLASS_CAMERASETTINGSSTATE byte = 15
	ARCOMMANDS_ID_COMMON_CLASS_GPS                 byte = 16
	ARCOMMANDS_ID_COMMON_CLASS_FLIGHTPLANSTATE     byte = 17
	ARCOMMANDS_ID_COMMON_CLASS_FLIGHTPLANEVENT     byte = 19
	ARCOMMANDS_ID_COMMON_CLASS_ARLIBSVERSIONSSTATE byte = 18

	// eARCOMMANDS_ID_ARDRONE3_PILOTING_CMD
	ARCOMMANDS_ID_ARDRONE3_PILOTING_CMD_FLATTRIM        byte = 0
	ARCOMMANDS_ID_ARDRONE3_PILOTING_CMD_TAKEOFF         byte = 1
	ARCOMMANDS_ID_ARDRONE3_PILOTING_CMD_PCMD            byte = 2
	ARCOMMANDS_ID_ARDRONE3_PILOTING_CMD_LANDING         byte = 3
	ARCOMMANDS_ID_ARDRONE3_PILOTING_CMD_EMERGENCY       byte = 4
	ARCOMMANDS_ID_ARDRONE3_PILOTING_CMD_NAVIGATEHOME    byte = 5
	ARCOMMANDS_ID_ARDRONE3_PILOTING_CMD_AUTOTAKEOFFMODE byte = 6
	ARCOMMANDS_ID_ARDRONE3_PILOTING_CMD_MAX             byte = 7

	// eARCOMMANDS_ID_ARDRONE3_MEDIARECORD_CMD
	ARCOMMANDS_ID_ARDRONE3_MEDIARECORD_CMD_PICTURE   byte = 0
	ARCOMMANDS_ID_ARDRONE3_MEDIARECORD_CMD_VIDEO     byte = 1
	ARCOMMANDS_ID_ARDRONE3_MEDIARECORD_CMD_PICTUREV2 byte = 2
	ARCOMMANDS_ID_ARDRONE3_MEDIARECORD_CMD_VIDEOV2   byte = 3
	ARCOMMANDS_ID_ARDRONE3_MEDIARECORD_CMD_MAX       byte = 4

	// eARCOMMANDS_ARDRONE3_MEDIARECORD_VIDEO_RECORD
	ARCOMMANDS_ARDRONE3_MEDIARECORD_VIDEO_RECORD_STOP  byte = 0
	ARCOMMANDS_ARDRONE3_MEDIARECORD_VIDEO_RECORD_START byte = 1
	ARCOMMANDS_ARDRONE3_MEDIARECORD_VIDEO_RECORD_MAX   byte = 2

	// eARCOMMANDS_ID_COMMON_COMMON_CMD
	ARCOMMANDS_ID_COMMON_COMMON_CMD_ALLSTATES   byte = 0
	ARCOMMANDS_ID_COMMON_COMMON_CMD_CURRENTDATE byte = 1
	ARCOMMANDS_ID_COMMON_COMMON_CMD_CURRENTTIME byte = 2
	ARCOMMANDS_ID_COMMON_COMMON_CMD_REBOOT      byte = 3
	ARCOMMANDS_ID_COMMON_COMMON_CMD_MAX         byte = 4

	// eARCOMMANDS_ID_COMMON_COMMONSTATE_CMD;
	ARCOMMANDS_ID_COMMON_COMMONSTATE_CMD_ALLSTATESCHANGED                    byte = 0
	ARCOMMANDS_ID_COMMON_COMMONSTATE_CMD_BATTERYSTATECHANGED                 byte = 1
	ARCOMMANDS_ID_COMMON_COMMONSTATE_CMD_MASSSTORAGESTATELISTCHANGED         byte = 2
	ARCOMMANDS_ID_COMMON_COMMONSTATE_CMD_MASSSTORAGEINFOSTATELISTCHANGED     byte = 3
	ARCOMMANDS_ID_COMMON_COMMONSTATE_CMD_CURRENTDATECHANGED                  byte = 4
	ARCOMMANDS_ID_COMMON_COMMONSTATE_CMD_CURRENTTIMECHANGED                  byte = 5
	ARCOMMANDS_ID_COMMON_COMMONSTATE_CMD_MASSSTORAGEINFOREMAININGLISTCHANGED byte = 6
	ARCOMMANDS_ID_COMMON_COMMONSTATE_CMD_WIFISIGNALCHANGED                   byte = 6
	ARCOMMANDS_ID_COMMON_COMMONSTATE_CMD_SENSORSSTATESLISTCHANGED            byte = 7
	ARCOMMANDS_ID_COMMON_COMMONSTATE_CMD_MAX                                 byte = 8

	// eARMEDIA_ENCAPSULER_CODEC
	CODEC_UNKNNOWN     byte = 0
	CODEC_VLIB         byte = 1
	CODEC_P264         byte = 2
	CODEC_MPEG4_VISUAL byte = 3
	CODEC_MPEG4_AVC    byte = 4
	CODEC_MOTION_JPEG  byte = 5

	// eARMEDIA_ENCAPSULER_FRAME_TYPE;
	ARMEDIA_ENCAPSULER_FRAME_TYPE_UNKNNOWN byte = 0
	ARMEDIA_ENCAPSULER_FRAME_TYPE_I_FRAME  byte = 1
	ARMEDIA_ENCAPSULER_FRAME_TYPE_P_FRAME  byte = 2
	ARMEDIA_ENCAPSULER_FRAME_TYPE_JPEG     byte = 3
	ARMEDIA_ENCAPSULER_FRAME_TYPE_MAX      byte = 4

	// eARNETWORK_MANAGER_INTERNAL_BUFFER_ID
	ARNETWORK_MANAGER_INTERNAL_BUFFER_ID_PING byte = 0
	ARNETWORK_MANAGER_INTERNAL_BUFFER_ID_PONG byte = 1
	ARNETWORK_MANAGER_INTERNAL_BUFFER_ID_MAX  byte = 3

	// eARNETWORKAL_FRAME_TYPE
	ARNETWORKAL_FRAME_TYPE_UNINITIALIZED    byte = 0
	ARNETWORKAL_FRAME_TYPE_ACK              byte = 1
	ARNETWORKAL_FRAME_TYPE_DATA             byte = 2
	ARNETWORKAL_FRAME_TYPE_DATA_LOW_LATENCY byte = 3
	ARNETWORKAL_FRAME_TYPE_DATA_WITH_ACK    byte = 4
	ARNETWORKAL_FRAME_TYPE_MAX              byte = 5

	ARCOMMANDS_ID_ARDRONE3_SPEEDSETTINGS_CMD_MAXVERTICALSPEED byte = 0
	ARCOMMANDS_ID_ARDRONE3_SPEEDSETTINGS_CMD_MAXROTATIONSPEED byte = 1
	ARCOMMANDS_ID_ARDRONE3_SPEEDSETTINGS_CMD_HULLPROTECTION byte = 2
	ARCOMMANDS_ID_ARDRONE3_SPEEDSETTINGS_CMD_OUTDOOR byte = 3

	ARCOMMANDS_ID_ARDRONE3_MEDIASTREAMING_CMD_VIDEOENABLE byte = 0
	ARCOMMANDS_ID_ARDRONE3_MEDIASTREAMING_CMD_VIDEOSTREAMMODE byte = 1
)
