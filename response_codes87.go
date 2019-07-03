package iso8583

const (
	//Sources:
	//https://help.sap.com/doc/saphelp_pos22/2.2/en-US/e8/097fdfd4164e639361c7cbeaf306f7/content.htm?no_cache=true
	//https://www.admfactory.com/iso8583-flows-data-elements-meaning-and-values/

	RC87_APPROVED                                                 = "00" //Successful approval/completion or that V.I.P. PIN verification is valid
	RC87_REFER_TO_ISSUER                                          = "01" //Refer to card issuer
	RC87_REFER_TO_ISSUER_SPECIAL_CONDITION                        = "02" //Refer to card issuer, special condition
	RC87_INVALID_MERCHANT                                         = "03" //Invalid merchant or service provider
	RC87_PICKUP_CARD                                              = "04" //Pickup card
	RC87_DO_NOT_HONOR                                             = "05" //Do not honor
	RC87_ERROR                                                    = "06" //Error
	RC87_PICKUP_CARD_SPECIAL_CONDITION                            = "07" //Pickup card, special condition (other than lost/stolen card)
	RC87_PARTIAL_APPROVAL                                         = "10" //Partial Approval
	RC87_VIP_APPROVAL                                             = "51" //V.I.P. approval
	RC87_INVALID_TXN                                              = "12" //Invalid transaction
	RC87_INVALID_AMOUNT                                           = "13" //Invalid amount (currency conversion field overflow)
	RC87_INVALID_ACCOUNT_NUMBER                                   = "14" //Invalid account number (no such number)
	RC87_NO_SUCH_ISSUER                                           = "15" //No such issuer
	RC87_CUSTOMER_CANCELLATION                                    = "17" //Customer cancellation
	RC87_RE_ENTER_TXN                                             = "19" //Re-enter transaction
	RC87_INVALID_RESPONSE                                         = "20" //Invalid response
	RC87_NO_ACTION_TAKEN                                          = "21" //No action taken (unable to back out prior transaction)
	RC87_SUSPECTED_MALFUNCTION                                    = "22" //Suspected Malfunction
	RC87_UNABLE_TO_LOCATE_RECORD                                  = "25" //Unable to locate record in file, or account number is missing from the inquiry
	RC87_FILE_TEMPORARILY_UNAVAILABLE                             = "28" //File is temporarily unavailable
	RC87_FORMAT_ERROR                                             = "30" //Format Error
	RC87_PICKUP_CARD_LOST                                         = "41" //Pickup card (lost card)
	RC87_PICKUP_CARD_STOLEN                                       = "43" //Pickup card (stolen card)
	RC87_INSUFFICIENT_FUNDS                                       = "51" //Insufficient funds
	RC87_NO_CHECKING_ACCOUNT                                      = "52" //No checking account
	RC87_NO_SAVINGS_ACCOUNT                                       = "53" //No savings account
	RC87_EXPIRED_CARD                                             = "54" //Expired card
	RC87_INCORRECT_PIN                                            = "55" //Incorrect PIN
	RC87_TXN_NOT_PERMITTED_TO_CARDHOLDER                          = "57" //Transaction not permitted to cardholder
	RC87_TXN_NOT_ALLOWED_AT_TERMINAL                              = "58" //Transaction not allowed at terminal
	RC87_SUSPECTED_FRAUD                                          = "59" //Suspected fraud
	RC87_ACTIVITY_AMOUNT_LIMIT_EXCEEDED                           = "61" //Activity amount limit exceeded
	RC87_RESTRICTED_CARD                                          = "62" //Restricted card (for example, in Country Exclusion table)
	RC87_SECURITY_VIOLATION                                       = "63" //Security violation
	RC87_ACTIVITY_COUNT_LIMIT_EXCEEDED                            = "65" //Activity count limit exceeded
	RC87_RESPONSE_RECEIVED_TOO_LATE                               = "68" //Response received too late
	RC87_PIN_ENTRY_TRIES_EXCEEDED                                 = "75" //Allowable number of PIN-entry tries exceeded
	RC87_UNABLE_TO_LOCATE_PREVIOUS_MESSAGE                        = "76" //Unable to locate previous message (no match on Retrieval Reference number)
	RC87_INCONSISTENT_REPEAT_REVERSAL_DATA                        = "77" //Previous message located for a repeat or reversal, but repeat or reversal data are inconsistent with original message
	RC87_BLOCKED_FIRST_USE                                        = "78" //'Blocked, first used'-The transaction is from a new cardholder, and the card has not been properly unblocked.
	RC87_VISA_TXN_CREDIT_ISSUER_UNAVAILABLE                       = "80" //Visa transactions: credit issuer unavailable. Private label and check acceptance: Invalid date
	RC87_PIN_CRYPTOGRAPHIC_ERROR                                  = "81" //PIN cryptographic error found (error found by VIC security module during PIN decryption)
	RC87_NEGATIVE_CAM_DCVV_ICVV_OR_CVV_RESULT                     = "82" //Negative CAM, dCVV, iCVV, or CVV results
	RC87_UNABLE_TO_VERIFY_PIN                                     = "83" //Unable to verify PIN
	RC87_NO_REASON_TO_DECLINE_A_REQUEST                           = "85" //No reason to decline a request for account number verification, address verification, CVV2 verification, or a credit voucher or merchandise return
	RC87_ISSUER_OR_SWITCH_UNAVAILABLE                             = "91" //Issuer unavailable or switch inoperative (STIP not applicable or available for this transaction)
	RC87_ROUTING_DESTINATION_NOT_FOUND                            = "92" //Destination cannot be found for routing
	RC87_ILLEGAL_TXN                                              = "93" //Transaction cannot be completed, violation of law
	RC87_DUPLICATE_TRANSMISSION                                   = "94" //Duplicate Transmission
	RC87_RECONCILIATION_ERROR                                     = "95" //Reconcile error
	RC87_SYSTEM_MALFUNCTION_OR_FIELD_ERROR                        = "96" //System malfunction, System malfunction or certain field error conditions
	RC87_US_ACQUIRER_SURCHARGE_AMOUNT_NOT_PERMITTED_ON_VISA_CARDS = "B1" //Surcharge amount not permitted on Visa cards (U.S. acquirers only)
	RC87_FORCE_STIP                                               = "N0" //Force STIP
	RC87_CASH_SERVICE_NOT_AVAILABLE                               = "N3" //Cash service not available
	RC87_CASHBACK_REQUEST_EXCEEDS_ISSUER_LIMIT                    = "N4" //Cashback request exceeds issuer limit
	RC87_DECLINE_FOR_CVV2_FAILURE                                 = "N7" //Decline for CVV2 failure
	RC87_INVALID_BILLER_INFORMATION                               = "P2" //Invalid biller information
	RC87_PIN_CHANGE_UNBLOCK_REQUEST_DECLINED                      = "P5" //PIN Change/Unblock request declined
	RC87_UNSAFE_PIN                                               = "P6" //Unsafe PIN
	RC87_CARD_AUTHENTICATION_FAILED                               = "Q1" //Card Authentication failed
	RC87_STOP_PAYMENT_ORDER                                       = "R0" //Stop Payment Order
	RC87_REVOCATION_OF_AUTHORIZATION_ORDER                        = "R1" //Revocation of Authorization Order
	RC87_REVOCATION_OF_ALL_AUTHORIZATIONS_ORDER                   = "R3" //Revocation of All Authorizations Order
	RC87_FORWARD_TO_ISSUER1                                       = "XA" //Forward to issuer
	RC87_FORWARD_TO_ISSUER2                                       = "XD" //Forward to issuer
	RC87_UNABLE_TO_GO_ONLINE                                      = "Z3" //Unable to go online
)
