package iso8583

const (
	//Source: https://help.sap.com/doc/saphelp_pos22/2.2/en-US/e8/097fdfd4164e639361c7cbeaf306f7/content.htm?no_cache=true

	//ISO message definitions

	// authorization messages
	AUTH_REQ           = "100" // authorization request
	AUTH_REQ_REPEAT    = "101" // authorization request repeat
	AUTH_REQ_RESP      = "110" // authorization request response
	AUTH_ADVICE        = "120" // authorization advice
	AUTH_ADVICE_REPEAT = "121" // authorization advice repeat
	AUTH_ADVICE_RESP   = "130" // authorization advice response
	AUTH_NOTIFY        = "140" // authorization notification
	AUTH_NOTIFY_RESP   = "150" // authorization notification response

	// verification messages
	VERIF_REQ           = "104" // verification request
	VERIF_REQ_REPEAT    = "105" // verification request repeat
	VERIF_REQ_RESP      = "114" // verification request response
	VERIF_ADVICE        = "124" // verification advice
	VERIF_ADVICE_REPEAT = "125" // verification advice repeat
	VERIF_ADVICE_RESP   = "134" // verification advice response
	VERIF_NOTIFY        = "144" // verification notification
	VERIF_NOTIFY_RESP   = "154" // verification notification response

	// financial messages
	FIN_REQ           = "200" // financial request
	FIN_REQ_REPEAT    = "201" // financial request repeat
	FIN_REQ_RESP      = "210" // financial request response
	FIN_ADVICE        = "220" // financial advice
	FIN_ADVICE_REPEAT = "221" // financial advice repeat
	FIN_ADVICE_RESP   = "230" // financial advice response
	FIN_NOTIFY        = "240" // financial notification
	FIN_NOTIFY_RESP   = "250" // financial notification response

	// file action messages
	FILE_ACTION_REQ           = "304" // file action request
	FILE_ACTION_REQ_REPEAT    = "305" // file action request repeat
	FILE_ACTION_REQ_RESP      = "314" // file action request response
	FILE_ACTION_ADVICE        = "324" // file action advice
	FILE_ACTION_ADVICE_REPEAT = "325" // file action advice repeat
	FILE_ACTION_ADVICE_RESP   = "334" // file action advice response
	FILE_ACTION_NOTIFY        = "344" // file action notification
	FILE_ACTION_NOTIFY_RESP   = "354" // file action notification response

	// reversal/chargeback messages
	REV_ADVICE           = "420" // reversal advice
	REV_ADVICE_REPEAT    = "421" // reversal advice repeat
	REV_ADVICE_RESP      = "430" // reversal advice response
	REV_ADVICE_NOTIFY    = "440" // reversal notification
	CHARGE_ADVICE        = "422" // chargeback advice
	CHARGE_ADVICE_REPEAT = "423" // chargeback advice repeat
	CHARGE_ADVICE_RESP   = "432" // chargeback advice response
	CHARGE_NOTIFY        = "442" // chargeback advice notification
	CHARGE_NOTIFY_RESP   = "452" // chargeback advice notification response

	// acquirer reconciliation messages
	ACQ_RECON_REQ           = "500" // acquirer reconciliation request
	ACQ_RECON_REQ_REPEAT    = "501" // acquirer reconciliation request repeat
	ACQ_RECON_REQ_RESP      = "510" // acquirer reconciliation request response
	ACQ_RECON_ADVICE        = "520" // acquirer reconciliation advice
	ACQ_RECON_ADVICE_REPEAT = "521" // acquirer reconciliation advice repeat
	ACQ_RECON_ADVICE_RESP   = "530" // acquirer reconciliation advice response
	ACQ_RECON_NOTIFY        = "540" // acquirer reconciliation notification
	ACQ_RECON_NOTIFY_RESP   = "550" // acquirer reconciliation notification response

	// issuer reconciliation messages
	ISSR_RECON_REQ           = "502" // issuer reconciliation request
	ISSR_RECON_REQ_REPEAT    = "503" // issuer reconciliation request repeat
	ISSR_RECON_REQ_RESP      = "512" // issuer reconciliation request response
	ISSR_RECON_ADVICE        = "522" // issuer reconciliation advice
	ISSR_RECON_ADVICE_REPEAT = "523" // issuer reconciliation advice repeat
	ISSR_RECON_ADVICE_RESP   = "532" // issuer reconciliation afvice response
	ISSR_RECON_NOTIFY        = "542" // issuer reconciliation notification
	ISSR_RECON_NOTIFY_RESP   = "552" // issuer reconciliation notification response

	// network administrative messages
	ADMIN_REQ        = "604" // network administrative request
	ADMIN_REQ_REPEAT = "605" // network administrative request repeat
	ADMIN_REQ_RESP   = "614" // network administrative request response

	ADMIN_ADVICE        = "624" // administrative advice
	ADMIN_ADVICE_REPEAT = "625" // administrative advice repeat
	ADMIN_ADVICE_RESP   = "634" // administrative advice response

	ADMIN_NOTIFY      = "644" // administrative notification
	ADMIN_NOTIFY_RESP = "654" // administrative notification response

	// fee collection messages
	ACQ_FEE_ADVICE        = "720" // acquirer fee collection advice
	ACQ_FEE_ADVICE_REPEAT = "721" // acquirer fee collection advice repeat
	ACQ_FEE_ADVICE_RESP   = "730" // acquirer fee collection advice response
	ACQ_FEE_NOTIFY        = "740" // acquirer fee collection notification
	ACQ_FEE_NOTIFY_RESP   = "750" // acquirer fee collection notification response

	ISSR_FEE_ADVICE        = "722" // issuer fee collection advice
	ISSR_FEE_ADVICE_REPEAT = "723" // issuer fee collection advice repeat
	ISSR_FEE_ADVICE_RESP   = "732" // issuer fee collection advice response
	ISSR_FEE_NOTIFY        = "742" // issuer fee collection notification
	ISSR_FEE_NOTIFY_RESP   = "752" // issuer fee collection notification response

	// network management messages
	NET_MGMT_REQ           = "804" // network management request
	NET_MGMT_REQ_REPEAT    = "805" // network management request repeat
	NET_MGMT_REQ_RESP      = "814" // network management request response
	NET_MGMT_ADVICE        = "824" // network management advice
	NET_NGMT_ADVICE_REPEAT = "825" // network management advice repeat
	NET_MGMT_ADVICE_RESP   = "834" // network management advice response
	NET_MGMT_NOTIFY        = "844" // network management notification
	NET_MGMT_NOTIFY_RESP   = "854" // network management notification resp

	//ISO 8583(1993) action code - field 39

	//Authorization and transaction messages: 110,120,121,140 and 210,220,221 and 240
	//Meaning: indicate that the transaction has been approved
	RC_APPROVED              = "000" // approved
	RC_APPROVED_WITH_ID      = "001" // honor with identification
	RC_APPROVED_PARTIAL      = "002" // approved for partial amount
	RC_APPROVED_VIP          = "003" // approved(VIP)
	RC_APPROVED_TRACK3       = "004" // approved; update track 3
	RC_APPROVED_ACCT_SPEC    = "005" // approved, account type specified by card issuer
	RC_APPROVED_PARTIAL_SPEC = "006" // approved for partial amount; account type specified by card issuer
	RC_APPROVED_ICC          = "007" // approved, update ICC
	RC_APPROVED_NEED_CNFM    = "80"  // approved, but need confirmation(used for CIBC and NOVA debit card processing mode

	//Authorization and transaction messages: 110,120,121,140 and 210,220,221 and 240
	//Meaning: indicate that the transaction has processed by or on behalf of issuer and denied (not requiring a card pick-up)
	RC_DECLINED_DO_NOT_HONOR         = "00" // do not honor
	RC_DECLINED_EXPIRED_CARD         = "01" // expired card
	RC_DECLINED_SUSPECTED            = "02" // suspected fraud
	RC_DECLINED_CONTACT_ACQ          = "03" // card acceptor contact acquirer
	RC_DECLINED_RESTRICTED           = "04" // restricted card
	RC_DECLINED_CALL_ACQ             = "05" // card acceptor call acquirer's security department
	RC_DECLINED_PIN_EXCEED           = "06" // allowable PIN tries exceeded
	RC_DECLINED_REFER_ISSR           = "07" // refer to card issuer
	RC_DECLINED_REFER_ISSR_COND      = "08" // refer to card issuer's special conditions
	RC_DECLINED_INVALID_MERCHANT     = "09" // invalid merchant
	RC_DECLINED_INVALID_AMOUNT       = "10" // invalid amount
	RC_DECLINED_INVALID_CARD         = "11" // invalid card number
	RC_DECLINED_PIN_REQUIRED         = "12" // PIN data required
	RC_DECLINED_UNACCEPT_FEE         = "13" // unacceptable fee
	RC_DECLINED_ACCT_REQ             = "14" // no account of type requested
	RC_DECLINED_FUNC_NOT_SUPPORT     = "15" // requested function not supported
	RC_DECLINED_NOT_SUFF_FUNDS       = "16" // not sufficient funds
	RC_DECLINED_INCORRECT_PIN        = "17" // incorrect PIN
	RC_DECLINED_NO_CARD_RECORD       = "18" // no card record
	RC_DECLINED_NOT_ALLOW_CARDHOLDER = "19" // transaction not permitted to cardholder

	RC_DECLINED_NOT_ALLOW_TERMINAL  = "20" // transaction not permitted to terminal
	RC_DECLINED_EXCEED_AMOUNT_LIMIT = "21" // exceeds withdrawal amount limit
	RC_DECLINED_VIOLATION           = "22" // security violation
	RC_DECLINED_EXCEED_FREQ_LIMIT   = "23" // exceeds withdrawal frequency limit
	RC_DECLINED_VIOLATION_LAW       = "24" // violation of law
	RC_DECLINED_NOT_EFFECTIVE       = "25" // card not effective
	RC_DECLINED_INVALID_PIN         = "26" // invalid PIN block
	RC_DECLINED_PIN_LENGTH          = "27" // PIN length error
	RC_DECLINED_PIN_SYNC            = "28" // PIN key synch error
	RC_DECLINED_SUSPECTED_COUNTER   = "29" // suspected counterfeit card

	//Authorization and transaction messages: 110,120,121,140 and 210,220,221 and 240
	//Meaning: indicate that the transaction has processed by or on behalf of issuer and denied (requiring a card to be pick-up.)
	RC_PICKUP_DO_NOT_HONOR      = "200" // do not honor
	RC_PICKUP_EXPIRED_CARD      = "201" // expired card
	RC_PICKUP_SUSPECTED         = "202" // suspected fraud
	RC_PICKUP_CONTACT_ACQ       = "203" // card acceptor contact acquirer
	RC_PICKUP_RESTRICTED        = "204" // restricted card
	RC_PICKUP_CALL_ACQ          = "205" // card acceptor call acquirer's security department
	RC_PICKUP_PIN_EXCEED        = "206" // allowable PIN tries exceeded
	RC_PICKUP_SPEC_COND         = "207" // special condition
	RC_PICKUP_LOST              = "208" // lost card
	RC_PICKUP_STOLEN            = "209" // stolen card
	RC_PICKUP_SUSPECTED_COUNTER = "210" // suspected countfeid card

	//File action messages: 314 324,325 and 344
	//Meaning: indicates result of file action
	RC_FILE_SUCCESS              = "300" // successful
	RC_FILE_NOT_SUPPORT          = "301" // not supported by receiver
	RC_FILE_UNABLE_LOCATE_RECORD = "302" // unable to locate record on file
	RC_FILE_DUP_REPLACE          = "303" // duplicate record; old record replaced
	RC_FILE_EDIT_ERROR           = "304" // field edit error
	RC_FILE_LOCKED_OUT           = "305" // file locked out
	RC_FILE_NOT_SUCCESS          = "306" // not successful
	RC_FILE_FORMAT_ERROR         = "307" // format error
	RC_FILE_DUP_REJECT           = "308" // duplicate; new record rejected
	RC_FILE_UNKNOWN              = "309" // unknown file

	//Reversal and chargeback messages: 430,432,440 and 442
	//Meaning: result of the reversal or chargeback.
	RC_REVERSAL_ACCEPT = "400" // accepted

	//Reconciliation messages: 510,512,530 and 532
	//Meaning: result of reconciliation
	RC_RECON_IN_BALANCE          = "500" // reconciled; in balance
	RC_RECON_OUT_BALANCE         = "501" // reconciled; out balance
	RC_RECON_AMOUNT_NOT_RECON    = "502" // amount not reconciled; total provided
	RC_RECON_TOTAL_NOT_AVAILABLE = "503" // totals not avalable
	RC_RECON_NOT_RECON           = "504" // not reconciled; totals provided

	//Adminitrative request messages: 614;624;625 and 644
	//Meaning: result of operation
	RC_ADMIN_ACCEPT              = "600" // accepted
	RC_ADMIN_NOT_TRACE_ORIGIN    = "601" // not able to trace back original transaction
	RC_ADMIN_INVALID_REFERENCE   = "602" // invalid reference number
	RC_ADMIN_PAN_INCOMPATIBLE    = "603" // reference number/PAN incompatible
	RC_ADMIN_PHOTO_NOT_AVAILABLE = "604" // POS photograph is not available
	RC_ADMIN_ITEM_SUPP           = "605" // item supplied
	RC_ADMIN_DOC_NOT_SUPP        = "606" // request cannot be fulfilled-required/requested documentation is not available

	//Fee collection messages: 720;721;740;722;723 and 742
	RC_FEE_ACCEPT = "700" // accepted

	//Network management messages:  814;824;825 and 844
	RC_NETWORK_ACCEPT       = "800" // accepted
	RC_NETWORK_NO_LIABILITY = "900" // advice acknowledged; no financial liability accepted
	RC_NETWORK_LIABILITY    = "901" // advice acknowledged; financial liability accepted

	//Transaction request and response messages
	//Meanint: indicate transaction could not be processed
	RC_REJECT_INVALID_TXN              = "902" // invalid transaction
	RC_REJECT_RE_ENTER_TXN             = "903" // re-enter transaction
	RC_REJECT_FORMAT_ERROR             = "904" // format error
	RC_REJECT_ACQ_NOT_SUPP             = "905" // acquirer not supported by switch
	RC_REJECT_CUTOVER_IN_PROCESS       = "906" // cutover in process
	RC_REJECT_ISSUER_INOPERATIVE       = "907" // card issuer or switch inoperative
	RC_REJECT_DEST_NOT_FOUND           = "908" // transaction destination cannot be found for routing
	RC_REJECT_SYSTEM_MALFUNCTION       = "909" // system malfunction
	RC_REJECT_ISSUER_SIGNOFF           = "910" // card issuer signed off
	RC_REJECT_ISSUER_TIMEOUT           = "911" // card issuer timed out
	RC_REJECT_ISSUER_NOT_AVAILABLE     = "912" // card issuer unavailable
	RC_REJECT_DUP_TRANSMISSION         = "913" // duplicate transmission
	RC_REJECT_NOT_TRACE_ORIGIN         = "914" // not able to trace back to original transaction
	RC_REJECT_CHECKPOINT_ERROR         = "915" // reconciliation cutover or checkpoint error
	RC_REJECT_MAC_ERROR                = "916" // MAC incorrect
	RC_REJECT_MAC_KEY_SYNC             = "917" // MAC key sync error
	RC_REJECT_NO_COMM_KEY              = "918" // no communication keys available for use
	RC_REJECT_ENCRYPTION_KEY_SYNC      = "919" // encryption key sync error
	RC_REJECT_SECURITY_ERROR_TRY_AGAIN = "920" // security software/hardware error - try again
	RC_REJECT_SECURITY_ERROR_NO_ACTION = "921" // security software/hardware error - no action
	RC_REJECT_MSGNO_ERROR              = "922" // message number out of sequence
	RC_REJECT_REQ_IN_PROCESS           = "923" // request in progress

	// 950-999 Used in advice response(xx3x) to indicate the reason for rejection of the transfer of financial liability.
	RC_REJECT_VIOLATION = "950" // violation of business arrangement

	// ISO 8583(1993) function code - field 24

	// 000-099 reserved for ISO use

	// 100-199 Used in 100;101;120;121 and 140 messages
	FUNC_AUTH_ORIGAUTH_ACCUAMT = "00" // original authorization - amount accurate
	FUNC_AUTH_ORIGAUTH_ESTIAMT = "01" // original authorization - amount estimated
	FUNC_AUTH_REPLAUTH_ACCUAMT = "02" // replacement authorization - amount accurate
	FUNC_AUTH_REPLAUTH_ESTIAMT = "03" // replacement authorization - amount estimated
	FUNC_AUTH_RESUBM_ACCUAMT   = "04" // resubmission - amount accurate
	FUNC_AUTH_RESUBM_ESTIAMT   = "05" // resubmission - amount estimated
	FUNC_AUTH_SUPMAUTH_ACCUAMT = "06" // supplementary authorization - amount accurate
	FUNC_AUTH_SUPMAUTH_ESTIAMT = "07" // supplementary authorization - amount estimated
	FUNC_AUTH_INQUIRY          = "08" // inquiry

	// 200-299 Used in 1200;1201,1220,1221 and 1240 message
	FUNC_FIN_ORIGFIN_TXN         = "200" // original financial request/advice
	FUNC_FIN_PREVAUTH_SAMEAMT    = "201" // previously approved authorization - amount same
	FUNC_FIN_PREVAUTH_DIFFAMT    = "202" // previously approved authorization - amount differs
	FUNC_FIN_RESUB_PREVFIN_DENY  = "203" // resubmission of a previously denied financial request
	FUNC_FIN_RESUB_PREVFIN_REVER = "204" // resubmission of a previously reversed financial transaction
	FUNC_FIN_FIRST_REPRE         = "205" // first representment
	FUNC_FIN_SECOND_REPRE        = "206" // second representment
	FUNC_FIN_THIRD_REPRE         = "207" // third or subsequent representment

	// 300-399 Used in 1304,1305,1324,1325 and 1344 messages
	FUNC_FILE_UNASSI     = "300" // unassigned
	FUNC_FILE_ADD_RECORD = "301" // addd record
	FUNC_FILE_CHG_RECORD = "302" // change record
	FUNC_FLLE_DEL_RECORD = "303" // delete record
	FUNC_FILE_REP_RECORD = "304" // replace record
	FUNC_FILE_INQ        = "305" // inquiry
	FUNC_FILE_REP_FILE   = "306" // replace file
	FUNC_FILE_ADD_FILE   = "307" // add file
	FUNC_FILE_DEL_FILE   = "308" // delete file
	FUNC_FILE_CARD_ADMIN = "309" // card administration

	// 400-449 Used in 1420,1421 and 1440 messages to indicate the function of the reversal
	FUNC_REV_FULL = "400" // full reversal; transaction did not complete as approved
	FUNC_REV_PART = "401" // partial reversal; transaction did not complete for full amount

	// 450-499 Used in 1422,1423 and 1442 messages to indicate the function of the chargeback
	FUNC_CHARG_FIRST_FULL   = "450" // first chargeback; full
	FUNC_CHARG_SECOND_FULL  = "451" // second chargeback; full
	FUNC_CHARG_THIRD_FULL   = "452" // third or subsequent chargeback; full
	FUNC_CHARG_FIRST_PART   = "453" // first chargeback; partial
	FUNC_CHARGE_SECOND_PART = "454" // second chargeback; partial
	FUNC_CHARGE_THIRD_PART  = "455" // third or subsequent; partial

	// 500-599 Used in 1500,1501,1502,1503,1520,1521,1522,1523,1540 and 1542 messages
	FUNC_RECON_FINAL          = "500" // final reconciliation
	FUNC_RECON_CHK_POINT      = "501" // checkpoint reconciliation
	FUNC_RECON_FINAL_CURR     = "502" // final reconciliation in a specified currency
	FUNC_RECON_CHK_POINT_CURR = "503" // checkpoint reconciliation in a specified currency
	FUNC_RECON_REQ            = "504" // request for reconciliation totals
	FUNC_RECON_SETTLE         = "570" // request to settle
	// 600-649 Used in 1604,1605,1624,1625 and 1644 messages for retrievals.
	FUNC_ADMIN_ORIGRECP_RETRREQ        = "600" // original receipt, retrieval request
	FUNC_ADMIN_ORIGRECP_RETRREQ_REPEAT = "601" // original receipt, repeat retrieval request
	FUNC_ADMIN_ORIGRECP_FULFILL        = "602" // original receipt, fulfillment
	FUNC_ADMIN_COPY_RETRREQ            = "603" // copy, retrieval request
	FUNC_ADMIN_COPY_RETRREQ_REPEAT     = "604" // copy, repeat retrieval request
	FUNC_ADMIN_COPY_FULFILL            = "605" // copyn; fulfillment
	FUNC_ADMIN_VEHICLE                 = "606" // vehicle rental agreement
	FUNC_ADMIN_HOTEL                   = "607" // hotel charge detail
	FUNC_ADMIN_POS                     = "608" // POS photograph
	FUNC_ADMIN_DEVY                    = "609" // proof of delivery
	FUNC_ADMIN_IMPRINT                 = "610" // imprint

	// 650-699 Used in 604,605,624,625 and 644 messages for administrative messages.
	FUNC_ADMIN_NOT_PARSE = "650" // unable to parse message

	// 700-799 Used in 720,721,740,722,723 and 742 messages.
	FUNC_FEE_COLL        = "700" // fee collection message
	FUNC_FEE_COLL_CANCEL = "701" // fee collection cancellation, full/partial

	// 800-899 Used in 804,805,824,825 and 844 messages.
	FUNC_NETWORK_SIGNON  = "801" // system condition/sign-on
	FUNC_NETWORK_SIGNOFF = "802" // system condition/sign-off
	FUNC_NETWORK_UNAVAIL = "803" // system condition/target system unavailable
	FUNC_NETWORK_BACKUP  = "804" // system condition/message originator's system in backup
	FUNC_NETWORK_SPECIAL = "805" // system condition/special instruction
	FUNC_NETWORK_ROUTE   = "806" // system condition/initate alternate routing

	FUNC_NETWORK_KEYCHANGE    = "811" // system security/key change
	FUNC_NETWORK_ALERT        = "812" // system security/security alert
	FUNC_NETWORK_PASSWDCHANGE = "813" // system security/password change
	FUNC_NETWORK_DEVICE_AUTH  = "814" // system security/device authentication

	FUNC_NETWORK_CUTOVER   = "821" // system accounting/cutover
	FUNC_NETWORK_CHK_POINT = "822" // system accounting/checkpoint

	FUNC_NETWORK_ECHO = "831" // system audit control/echo test

	// ISO 8583(1993) amount type codes - field ... ...
	// 00-19 account related balances
	// 00 reserved for ISO use
	AMT_LEDGER_BALANCE = "01" // account ledger balance
	AMT_AVAIL_BALANCE  = "02" // account available balance
	AMT_OWING          = "03" // account owning
	AMT_DUE            = "04" // account due
	AMT_AVAIL_CREDIT   = "05" // account available credit
	// 20-39 card related amounts
	AMT_REMAINING = "20" // amount remaining this cycle
	// 40-59 transaction related amounts
	AMT_CASH  = "40" // amount cash
	AMT_GOODS = "41" // amount goods and services

	// ISO 8583(1993) processing code - field 3(12)

	// 00-19 debits
	TPC_DB_GOODS_AND_SERVICE = "00" // goods and service
	TPC_DB_CASH              = "01" // cash
	TPC_DB_ADJUSTMENT        = "02" // adjustment
	TPC_DB_CHEQUE_GUAR       = "03" // cheque guarantee(funds guaranteed)
	TPC_DB_CHEQUE_VERI       = "04" // cheque verification(funds available but not quaranteed)
	TPC_DB_EURO_CHEQUE       = "05" // eurocheque
	TPC_DB_TRAVEL_CHEQUE     = "06" // traveller cheque
	TPC_DB_LETTER_CREDIT     = "07" // letter of credit
	TPC_DB_GIRO              = "8"  // giro(postal banking)
	TPC_DB_DISBURSE          = "9"  // goods and services with cash disbursement
	TPC_DB_NON_CASH          = "0"  // non-cash financial instrument(e.g.wire transfer)
	TPC_DB_QUASI             = "1"  // quasi-cash and scrip
	TPC_DB_SECOND_REQUEST    = "7"  // Tender Retail second request message

	// 20-29 credits
	TPC_CR_RETURN       = "20" // returns
	TPC_CR_DEPOSIT      = "21" // deposits
	TPC_CR_ADJUSTMENT   = "22" // adjustment
	TPC_CR_CHEQUE_GUAR  = "23" // cheque deposit guarantee
	TPC_CR_CHEQUE       = "24" // cheque deposit
	TPC_CASH_ADJUSTMENT = "28" // cash adjustment

	// 30-39 inquiry services
	TPC_FUND_INQUIRY    = "30" // available funds inquiry
	TPC_BALANCE_INQUIRY = "31" // balance inquiry

	// 40-49 transfer services
	TPC_TRANSFER            = "40" // cardholder accounts transfer
	TPC_TRANSFER_ADJUSTMENT = "48" // void cardholder accounts transfer

	// 50-59 payment services
	TPC_PAYMENT            = "50" // payment
	TPC_PAYMENT_ADJUSTMENT = "58" // payment adjustment

	// 90-99 reserved for private use
	TPC_ACTIVATE              = "90" // account activation
	TPC_ACTIVATE_ADJUSTMENT   = "91" // void account activation
	TPC_DEACTIVATE            = "92" // account deactivation
	TPC_DEACTIVATE_ADJUSTMENT = "93" // void account deactivation

	// ISO 8583 (1993) account type - field 3(34/56)
	ACCNT_DEFAULT      = "00" // default - unspecified
	ACCNT_SAVING       = "0"  // saving account - default
	ACCNT_CHEQUE       = "20" // cheque account - default
	ACCNT_CREDIT       = "30" // credit facility - default
	ACCNT_UNIVERSAL    = "40" // universal account - default
	ACCNT_INVESTMENT   = "50" // investment account - default
	ACCNT_PRIVATELABEL = "90" // private label card
	ACCNT_STOREVALUE   = "91" // store value card
	ACCNT_GIFT         = "92" // gift card

	// ISO Card Data Input Capability
	CARD_CAPTURE_UNKNOWN  = 0
	CARD_CAPTURE_MANUAL   = 1
	CARD_CAPTURE_SWIPED   = 2
	CARD_CAPTURE_SCANNED  = 3
	CARD_CAPTURE_OCR      = 4
	CARD_CAPTURE_ICC      = 5
	CARD_CAPTURE_KEYED    = 6
	CARD_CAPTURE_RESERVED = 7
)
