// WARNING: This file has automatically been generated on Thu, 23 Aug 2018 14:06:38 PDT.
// Code generated by https://git.io/c-for-go. DO NOT EDIT.

package cipvs

const (
	// VersionCode as defined in cipvs/ip_vs.h:12
	VersionCode = 0x010201
	// SvcFPersistent as defined in cipvs/ip_vs.h:21
	SvcFPersistent = 0x0001
	// SvcFHashed as defined in cipvs/ip_vs.h:22
	SvcFHashed = 0x0002
	// SvcFOnepacket as defined in cipvs/ip_vs.h:23
	SvcFOnepacket = 0x0004
	// SvcFSched1 as defined in cipvs/ip_vs.h:24
	SvcFSched1 = 0x0008
	// SvcFSched2 as defined in cipvs/ip_vs.h:25
	SvcFSched2 = 0x0010
	// SvcFSched3 as defined in cipvs/ip_vs.h:26
	SvcFSched3 = 0x0020
	// SvcFSchedShFallback as defined in cipvs/ip_vs.h:28
	SvcFSchedShFallback = SvcFSched1
	// SvcFSchedShPort as defined in cipvs/ip_vs.h:29
	SvcFSchedShPort = SvcFSched2
	// DestFAvailable as defined in cipvs/ip_vs.h:34
	DestFAvailable = 0x0001
	// DestFOverload as defined in cipvs/ip_vs.h:35
	DestFOverload = 0x0002
	// StateNone as defined in cipvs/ip_vs.h:40
	StateNone = 0x0000
	// StateMaster as defined in cipvs/ip_vs.h:41
	StateMaster = 0x0001
	// StateBackup as defined in cipvs/ip_vs.h:42
	StateBackup = 0x0002
	// ConnFFwdMask as defined in cipvs/ip_vs.h:82
	ConnFFwdMask = 0x0007
	// ConnFMasq as defined in cipvs/ip_vs.h:83
	ConnFMasq = 0x0000
	// ConnFLocalnode as defined in cipvs/ip_vs.h:84
	ConnFLocalnode = 0x0001
	// ConnFTunnel as defined in cipvs/ip_vs.h:85
	ConnFTunnel = 0x0002
	// ConnFDroute as defined in cipvs/ip_vs.h:86
	ConnFDroute = 0x0003
	// ConnFBypass as defined in cipvs/ip_vs.h:87
	ConnFBypass = 0x0004
	// ConnFSync as defined in cipvs/ip_vs.h:88
	ConnFSync = 0x0020
	// ConnFHashed as defined in cipvs/ip_vs.h:89
	ConnFHashed = 0x0040
	// ConnFNooutput as defined in cipvs/ip_vs.h:90
	ConnFNooutput = 0x0080
	// ConnFInactive as defined in cipvs/ip_vs.h:91
	ConnFInactive = 0x0100
	// ConnFOutSeq as defined in cipvs/ip_vs.h:92
	ConnFOutSeq = 0x0200
	// ConnFInSeq as defined in cipvs/ip_vs.h:93
	ConnFInSeq = 0x0400
	// ConnFSeqMask as defined in cipvs/ip_vs.h:94
	ConnFSeqMask = 0x0600
	// ConnFNoCport as defined in cipvs/ip_vs.h:95
	ConnFNoCport = 0x0800
	// ConnFTemplate as defined in cipvs/ip_vs.h:96
	ConnFTemplate = 0x1000
	// ConnFOnePacket as defined in cipvs/ip_vs.h:97
	ConnFOnePacket = 0x2000
	// ConnFBackupMask as defined in cipvs/ip_vs.h:100
	ConnFBackupMask = (ConnFFwdMask | ConnFNooutput | ConnFInactive | ConnFSeqMask | ConnFNoCport | ConnFTemplate)
	// ConnFBackupUpdMask as defined in cipvs/ip_vs.h:109
	ConnFBackupUpdMask = (ConnFInactive | ConnFSeqMask)
	// ConnFNfct as defined in cipvs/ip_vs.h:113
	ConnFNfct = (1 << 16)
	// ConnFDestMask as defined in cipvs/ip_vs.h:116
	ConnFDestMask = (ConnFFwdMask | ConnFOnePacket | ConnFNfct | 0)
	// SchednameMaxlen as defined in cipvs/ip_vs.h:121
	SchednameMaxlen = 16
	// PenameMaxlen as defined in cipvs/ip_vs.h:122
	PenameMaxlen = 16
	// IfnameMaxlen as defined in cipvs/ip_vs.h:123
	IfnameMaxlen = 16
	// PedataMaxlen as defined in cipvs/ip_vs.h:125
	PedataMaxlen = 255
	// GenlName as defined in cipvs/ip_vs.h:286
	GenlName = "IPVS"
	// GenlVersion as defined in cipvs/ip_vs.h:287
	GenlVersion = 0x1
	// CmdMax as defined in cipvs/ip_vs.h:324
	CmdMax = (__CmdMax - 1)
	// CmdAttrMax as defined in cipvs/ip_vs.h:338
	CmdAttrMax = (__CmdAttrMax - 1)
	// SvcAttrMax as defined in cipvs/ip_vs.h:367
	SvcAttrMax = (__SvcAttrMax - 1)
	// DestAttrMax as defined in cipvs/ip_vs.h:398
	DestAttrMax = (__DestAttrMax - 1)
	// DaemonAttrMax as defined in cipvs/ip_vs.h:418
	DaemonAttrMax = (__DaemonAttrMax - 1)
	// StatsAttrMax as defined in cipvs/ip_vs.h:443
	StatsAttrMax = (__StatsAttrMax - 1)
	// InfoAttrMax as defined in cipvs/ip_vs.h:453
	InfoAttrMax = (__InfoAttrMax - 1)
)

const (
	// CmdUnspec as declared in cipvs/ip_vs.h:296
	CmdUnspec = iota
	// CmdNewService as declared in cipvs/ip_vs.h:298
	CmdNewService = 1
	// CmdSetService as declared in cipvs/ip_vs.h:299
	CmdSetService = 2
	// CmdDelService as declared in cipvs/ip_vs.h:300
	CmdDelService = 3
	// CmdGetService as declared in cipvs/ip_vs.h:301
	CmdGetService = 4
	// CmdNewDest as declared in cipvs/ip_vs.h:303
	CmdNewDest = 5
	// CmdSetDest as declared in cipvs/ip_vs.h:304
	CmdSetDest = 6
	// CmdDelDest as declared in cipvs/ip_vs.h:305
	CmdDelDest = 7
	// CmdGetDest as declared in cipvs/ip_vs.h:306
	CmdGetDest = 8
	// CmdNewDaemon as declared in cipvs/ip_vs.h:308
	CmdNewDaemon = 9
	// CmdDelDaemon as declared in cipvs/ip_vs.h:309
	CmdDelDaemon = 10
	// CmdGetDaemon as declared in cipvs/ip_vs.h:310
	CmdGetDaemon = 11
	// CmdSetConfig as declared in cipvs/ip_vs.h:312
	CmdSetConfig = 12
	// CmdGetConfig as declared in cipvs/ip_vs.h:313
	CmdGetConfig = 13
	// CmdSetInfo as declared in cipvs/ip_vs.h:315
	CmdSetInfo = 14
	// CmdGetInfo as declared in cipvs/ip_vs.h:316
	CmdGetInfo = 15
	// CmdZero as declared in cipvs/ip_vs.h:318
	CmdZero = 16
	// CmdFlush as declared in cipvs/ip_vs.h:319
	CmdFlush = 17
	// __CmdMax as declared in cipvs/ip_vs.h:321
	__CmdMax = 18
)

const (
	// CmdAttrUnspec as declared in cipvs/ip_vs.h:328
	CmdAttrUnspec = iota
	// CmdAttrService as declared in cipvs/ip_vs.h:329
	CmdAttrService = 1
	// CmdAttrDest as declared in cipvs/ip_vs.h:330
	CmdAttrDest = 2
	// CmdAttrDaemon as declared in cipvs/ip_vs.h:331
	CmdAttrDaemon = 3
	// CmdAttrTimeoutTcp as declared in cipvs/ip_vs.h:332
	CmdAttrTimeoutTcp = 4
	// CmdAttrTimeoutTcpFin as declared in cipvs/ip_vs.h:333
	CmdAttrTimeoutTcpFin = 5
	// CmdAttrTimeoutUdp as declared in cipvs/ip_vs.h:334
	CmdAttrTimeoutUdp = 6
	// __CmdAttrMax as declared in cipvs/ip_vs.h:335
	__CmdAttrMax = 7
)

const (
	// SvcAttrUnspec as declared in cipvs/ip_vs.h:346
	SvcAttrUnspec = iota
	// SvcAttrAf as declared in cipvs/ip_vs.h:347
	SvcAttrAf = 1
	// SvcAttrProtocol as declared in cipvs/ip_vs.h:348
	SvcAttrProtocol = 2
	// SvcAttrAddr as declared in cipvs/ip_vs.h:349
	SvcAttrAddr = 3
	// SvcAttrPort as declared in cipvs/ip_vs.h:350
	SvcAttrPort = 4
	// SvcAttrFwmark as declared in cipvs/ip_vs.h:351
	SvcAttrFwmark = 5
	// SvcAttrSchedName as declared in cipvs/ip_vs.h:353
	SvcAttrSchedName = 6
	// SvcAttrFlags as declared in cipvs/ip_vs.h:354
	SvcAttrFlags = 7
	// SvcAttrTimeout as declared in cipvs/ip_vs.h:355
	SvcAttrTimeout = 8
	// SvcAttrNetmask as declared in cipvs/ip_vs.h:356
	SvcAttrNetmask = 9
	// SvcAttrStats as declared in cipvs/ip_vs.h:358
	SvcAttrStats = 10
	// SvcAttrPeName as declared in cipvs/ip_vs.h:360
	SvcAttrPeName = 11
	// SvcAttrStats64 as declared in cipvs/ip_vs.h:362
	SvcAttrStats64 = 12
	// __SvcAttrMax as declared in cipvs/ip_vs.h:364
	__SvcAttrMax = 13
)

const (
	// DestAttrUnspec as declared in cipvs/ip_vs.h:375
	DestAttrUnspec = iota
	// DestAttrAddr as declared in cipvs/ip_vs.h:376
	DestAttrAddr = 1
	// DestAttrPort as declared in cipvs/ip_vs.h:377
	DestAttrPort = 2
	// DestAttrFwdMethod as declared in cipvs/ip_vs.h:379
	DestAttrFwdMethod = 3
	// DestAttrWeight as declared in cipvs/ip_vs.h:380
	DestAttrWeight = 4
	// DestAttrUThresh as declared in cipvs/ip_vs.h:382
	DestAttrUThresh = 5
	// DestAttrLThresh as declared in cipvs/ip_vs.h:383
	DestAttrLThresh = 6
	// DestAttrActiveConns as declared in cipvs/ip_vs.h:385
	DestAttrActiveConns = 7
	// DestAttrInactConns as declared in cipvs/ip_vs.h:386
	DestAttrInactConns = 8
	// DestAttrPersistConns as declared in cipvs/ip_vs.h:387
	DestAttrPersistConns = 9
	// DestAttrStats as declared in cipvs/ip_vs.h:389
	DestAttrStats = 10
	// DestAttrAddrFamily as declared in cipvs/ip_vs.h:391
	DestAttrAddrFamily = 11
	// DestAttrStats64 as declared in cipvs/ip_vs.h:393
	DestAttrStats64 = 12
	// __DestAttrMax as declared in cipvs/ip_vs.h:395
	__DestAttrMax = 13
)

const (
	// DaemonAttrUnspec as declared in cipvs/ip_vs.h:406
	DaemonAttrUnspec = iota
	// DaemonAttrState as declared in cipvs/ip_vs.h:407
	DaemonAttrState = 1
	// DaemonAttrMcastIfn as declared in cipvs/ip_vs.h:408
	DaemonAttrMcastIfn = 2
	// DaemonAttrSyncId as declared in cipvs/ip_vs.h:409
	DaemonAttrSyncId = 3
	// DaemonAttrSyncMaxlen as declared in cipvs/ip_vs.h:410
	DaemonAttrSyncMaxlen = 4
	// DaemonAttrMcastGroup as declared in cipvs/ip_vs.h:411
	DaemonAttrMcastGroup = 5
	// DaemonAttrMcastGroup6 as declared in cipvs/ip_vs.h:412
	DaemonAttrMcastGroup6 = 6
	// DaemonAttrMcastPort as declared in cipvs/ip_vs.h:413
	DaemonAttrMcastPort = 7
	// DaemonAttrMcastTtl as declared in cipvs/ip_vs.h:414
	DaemonAttrMcastTtl = 8
	// __DaemonAttrMax as declared in cipvs/ip_vs.h:415
	__DaemonAttrMax = 9
)

const (
	// StatsAttrUnspec as declared in cipvs/ip_vs.h:427
	StatsAttrUnspec = iota
	// StatsAttrConns as declared in cipvs/ip_vs.h:428
	StatsAttrConns = 1
	// StatsAttrInpkts as declared in cipvs/ip_vs.h:429
	StatsAttrInpkts = 2
	// StatsAttrOutpkts as declared in cipvs/ip_vs.h:430
	StatsAttrOutpkts = 3
	// StatsAttrInbytes as declared in cipvs/ip_vs.h:431
	StatsAttrInbytes = 4
	// StatsAttrOutbytes as declared in cipvs/ip_vs.h:432
	StatsAttrOutbytes = 5
	// StatsAttrCps as declared in cipvs/ip_vs.h:434
	StatsAttrCps = 6
	// StatsAttrInpps as declared in cipvs/ip_vs.h:435
	StatsAttrInpps = 7
	// StatsAttrOutpps as declared in cipvs/ip_vs.h:436
	StatsAttrOutpps = 8
	// StatsAttrInbps as declared in cipvs/ip_vs.h:437
	StatsAttrInbps = 9
	// StatsAttrOutbps as declared in cipvs/ip_vs.h:438
	StatsAttrOutbps = 10
	// StatsAttrPad as declared in cipvs/ip_vs.h:439
	StatsAttrPad = 11
	// __StatsAttrMax as declared in cipvs/ip_vs.h:440
	__StatsAttrMax = 12
)

const (
	// InfoAttrUnspec as declared in cipvs/ip_vs.h:447
	InfoAttrUnspec = iota
	// InfoAttrVersion as declared in cipvs/ip_vs.h:448
	InfoAttrVersion = 1
	// InfoAttrConnTabSize as declared in cipvs/ip_vs.h:449
	InfoAttrConnTabSize = 2
	// __InfoAttrMax as declared in cipvs/ip_vs.h:450
	__InfoAttrMax = 3
)
