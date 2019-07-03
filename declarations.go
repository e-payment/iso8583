package iso8583

const (
	SPEC1987 = "1987"
	SPEC1993 = "1993"
	SPEC2003 = "2003"
)

var (
	specs = map[string]Spec{}

	//The maps below are used to perform bitwise operations on bitmaps, to detect if a certain bit is set or not
	//10000000, 01000000, 00100000, 00010000, 00001000, 00000100, 00000010, 00000001
	tA = [8]byte{128, 64, 32, 16, 8, 4, 2, 1}

	//01111111, 10111111, 11011111, 11101111, 11110111, 11111011, 11111101, 11111110
	tB = [8]byte{127, 191, 223, 239, 247, 251, 253, 254}

	Currencies = map[string]Currency{
		"104": Currency{
			Code:     "mmk",
			Country:  "MYANMAR",
			Name:     "Kyat",
			Decimals: 100,
		},
		"108": Currency{
			Code:     "bif",
			Country:  "BURUNDI",
			Name:     "Burundi Franc",
			Decimals: 1,
		},
		"116": Currency{
			Code:     "khr",
			Country:  "CAMBODIA",
			Name:     "Riel",
			Decimals: 100,
		},
		"124": Currency{
			Code:     "cad",
			Country:  "CANADA",
			Name:     "Canadian Dollar",
			Decimals: 100,
		},
		"132": Currency{
			Code:     "cve",
			Country:  "CABO VERDE",
			Name:     "Cabo Verde Escudo",
			Decimals: 100,
		},
		"136": Currency{
			Code:     "kyd",
			Country:  "CAYMAN ISLANDS (THE)",
			Name:     "Cayman Islands Dollar",
			Decimals: 100,
		},
		"144": Currency{
			Code:     "lkr",
			Country:  "SRI LANKA",
			Name:     "Sri Lanka Rupee",
			Decimals: 100,
		},
		"152": Currency{
			Code:     "clp",
			Country:  "CHILE",
			Name:     "Chilean Peso",
			Decimals: 1,
		},
		"156": Currency{
			Code:     "cny",
			Country:  "CHINA",
			Name:     "Yuan Renminbi",
			Decimals: 100,
		},
		"170": Currency{
			Code:     "cop",
			Country:  "COLOMBIA",
			Name:     "Colombian Peso",
			Decimals: 100,
		},
		"174": Currency{
			Code:     "kmf",
			Country:  "COMOROS (THE)",
			Name:     "Comorian Franc",
			Decimals: 1,
		},
		"188": Currency{
			Code:     "crc",
			Country:  "COSTA RICA",
			Name:     "Costa Rican Colon",
			Decimals: 100,
		},
		"191": Currency{
			Code:     "hrk",
			Country:  "CROATIA",
			Name:     "Kuna",
			Decimals: 100,
		},
		"192": Currency{
			Code:     "cup",
			Country:  "CUBA",
			Name:     "Cuban Peso",
			Decimals: 100,
		},
		"203": Currency{
			Code:     "czk",
			Country:  "CZECHIA",
			Name:     "Czech Koruna",
			Decimals: 100,
		},
		"208": Currency{
			Code:     "dkk",
			Country:  "GREENLAND",
			Name:     "Danish Krone",
			Decimals: 100,
		},
		"214": Currency{
			Code:     "dop",
			Country:  "DOMINICAN REPUBLIC (THE)",
			Name:     "Dominican Peso",
			Decimals: 100,
		},
		"222": Currency{
			Code:     "svc",
			Country:  "EL SALVADOR",
			Name:     "El Salvador Colon",
			Decimals: 100,
		},
		"230": Currency{
			Code:     "etb",
			Country:  "ETHIOPIA",
			Name:     "Ethiopian Birr",
			Decimals: 100,
		},
		"232": Currency{
			Code:     "ern",
			Country:  "ERITREA",
			Name:     "Nakfa",
			Decimals: 100,
		},
		"238": Currency{
			Code:     "fkp",
			Country:  "FALKLAND ISLANDS (THE) [MALVINAS]",
			Name:     "Falkland Islands Pound",
			Decimals: 100,
		},
		"242": Currency{
			Code:     "fjd",
			Country:  "FIJI",
			Name:     "Fiji Dollar",
			Decimals: 100,
		},
		"262": Currency{
			Code:     "djf",
			Country:  "DJIBOUTI",
			Name:     "Djibouti Franc",
			Decimals: 1,
		},
		"270": Currency{
			Code:     "gmd",
			Country:  "GAMBIA (THE)",
			Name:     "Dalasi",
			Decimals: 100,
		},
		"292": Currency{
			Code:     "gip",
			Country:  "GIBRALTAR",
			Name:     "Gibraltar Pound",
			Decimals: 100,
		},
		"320": Currency{
			Code:     "gtq",
			Country:  "GUATEMALA",
			Name:     "Quetzal",
			Decimals: 100,
		},
		"324": Currency{
			Code:     "gnf",
			Country:  "GUINEA",
			Name:     "Guinean Franc",
			Decimals: 1,
		},
		"328": Currency{
			Code:     "gyd",
			Country:  "GUYANA",
			Name:     "Guyana Dollar",
			Decimals: 100,
		},
		"332": Currency{
			Code:     "htg",
			Country:  "HAITI",
			Name:     "Gourde",
			Decimals: 100,
		},
		"340": Currency{
			Code:     "hnl",
			Country:  "HONDURAS",
			Name:     "Lempira",
			Decimals: 100,
		},
		"344": Currency{
			Code:     "hkd",
			Country:  "HONG KONG",
			Name:     "Hong Kong Dollar",
			Decimals: 100,
		},
		"348": Currency{
			Code:     "huf",
			Country:  "HUNGARY",
			Name:     "Forint",
			Decimals: 100,
		},
		"352": Currency{
			Code:     "isk",
			Country:  "ICELAND",
			Name:     "Iceland Krona",
			Decimals: 1,
		},
		"356": Currency{
			Code:     "inr",
			Country:  "INDIA",
			Name:     "Indian Rupee",
			Decimals: 100,
		},
		"360": Currency{
			Code:     "idr",
			Country:  "INDONESIA",
			Name:     "Rupiah",
			Decimals: 100,
		},
		"364": Currency{
			Code:     "irr",
			Country:  "IRAN (ISLAMIC REPUBLIC OF)",
			Name:     "Iranian Rial",
			Decimals: 100,
		},
		"368": Currency{
			Code:     "iqd",
			Country:  "IRAQ",
			Name:     "Iraqi Dinar",
			Decimals: 1000,
		},
		"376": Currency{
			Code:     "ils",
			Country:  "ISRAEL",
			Name:     "New Israeli Sheqel",
			Decimals: 100,
		},
		"388": Currency{
			Code:     "jmd",
			Country:  "JAMAICA",
			Name:     "Jamaican Dollar",
			Decimals: 100,
		},
		"392": Currency{
			Code:     "jpy",
			Country:  "JAPAN",
			Name:     "Yen",
			Decimals: 1,
		},
		"398": Currency{
			Code:     "kzt",
			Country:  "KAZAKHSTAN",
			Name:     "Tenge",
			Decimals: 100,
		},
		"400": Currency{
			Code:     "jod",
			Country:  "JORDAN",
			Name:     "Jordanian Dinar",
			Decimals: 1000,
		},
		"404": Currency{
			Code:     "kes",
			Country:  "KENYA",
			Name:     "Kenyan Shilling",
			Decimals: 100,
		},
		"408": Currency{
			Code:     "kpw",
			Country:  "KOREA (THE DEMOCRATIC PEOPLE’S REPUBLIC OF)",
			Name:     "North Korean Won",
			Decimals: 100,
		},
		"410": Currency{
			Code:     "krw",
			Country:  "KOREA (THE REPUBLIC OF)",
			Name:     "Won",
			Decimals: 1,
		},
		"414": Currency{
			Code:     "kwd",
			Country:  "KUWAIT",
			Name:     "Kuwaiti Dinar",
			Decimals: 1000,
		},
		"417": Currency{
			Code:     "kgs",
			Country:  "KYRGYZSTAN",
			Name:     "Som",
			Decimals: 100,
		},
		"418": Currency{
			Code:     "lak",
			Country:  "LAO PEOPLE’S DEMOCRATIC REPUBLIC (THE)",
			Name:     "Lao Kip",
			Decimals: 100,
		},
		"422": Currency{
			Code:     "lbp",
			Country:  "LEBANON",
			Name:     "Lebanese Pound",
			Decimals: 100,
		},
		"426": Currency{
			Code:     "lsl",
			Country:  "LESOTHO",
			Name:     "Loti",
			Decimals: 100,
		},
		"430": Currency{
			Code:     "lrd",
			Country:  "LIBERIA",
			Name:     "Liberian Dollar",
			Decimals: 100,
		},
		"434": Currency{
			Code:     "lyd",
			Country:  "LIBYA",
			Name:     "Libyan Dinar",
			Decimals: 1000,
		},
		"446": Currency{
			Code:     "mop",
			Country:  "MACAO",
			Name:     "Pataca",
			Decimals: 100,
		},
		"454": Currency{
			Code:     "mwk",
			Country:  "MALAWI",
			Name:     "Malawi Kwacha",
			Decimals: 100,
		},
		"458": Currency{
			Code:     "myr",
			Country:  "MALAYSIA",
			Name:     "Malaysian Ringgit",
			Decimals: 100,
		},
		"462": Currency{
			Code:     "mvr",
			Country:  "MALDIVES",
			Name:     "Rufiyaa",
			Decimals: 100,
		},
		"480": Currency{
			Code:     "mur",
			Country:  "MAURITIUS",
			Name:     "Mauritius Rupee",
			Decimals: 100,
		},
		"484": Currency{
			Code:     "mxn",
			Country:  "MEXICO",
			Name:     "Mexican Peso",
			Decimals: 100,
		},
		"496": Currency{
			Code:     "mnt",
			Country:  "MONGOLIA",
			Name:     "Tugrik",
			Decimals: 100,
		},
		"498": Currency{
			Code:     "mdl",
			Country:  "MOLDOVA (THE REPUBLIC OF)",
			Name:     "Moldovan Leu",
			Decimals: 100,
		},
		"504": Currency{
			Code:     "mad",
			Country:  "WESTERN SAHARA",
			Name:     "Moroccan Dirham",
			Decimals: 100,
		},
		"512": Currency{
			Code:     "omr",
			Country:  "OMAN",
			Name:     "Rial Omani",
			Decimals: 1000,
		},
		"516": Currency{
			Code:     "nad",
			Country:  "NAMIBIA",
			Name:     "Namibia Dollar",
			Decimals: 100,
		},
		"524": Currency{
			Code:     "npr",
			Country:  "NEPAL",
			Name:     "Nepalese Rupee",
			Decimals: 100,
		},
		"532": Currency{
			Code:     "ang",
			Country:  "SINT MAARTEN (DUTCH PART)",
			Name:     "Netherlands Antillean Guilder",
			Decimals: 100,
		},
		"533": Currency{
			Code:     "awg",
			Country:  "ARUBA",
			Name:     "Aruban Florin",
			Decimals: 100,
		},
		"548": Currency{
			Code:     "vuv",
			Country:  "VANUATU",
			Name:     "Vatu",
			Decimals: 1,
		},
		"554": Currency{
			Code:     "nzd",
			Country:  "TOKELAU",
			Name:     "New Zealand Dollar",
			Decimals: 100,
		},
		"558": Currency{
			Code:     "nio",
			Country:  "NICARAGUA",
			Name:     "Cordoba Oro",
			Decimals: 100,
		},
		"566": Currency{
			Code:     "ngn",
			Country:  "NIGERIA",
			Name:     "Naira",
			Decimals: 100,
		},
		"578": Currency{
			Code:     "nok",
			Country:  "SVALBARD AND JAN MAYEN",
			Name:     "Norwegian Krone",
			Decimals: 100,
		},
		"586": Currency{
			Code:     "pkr",
			Country:  "PAKISTAN",
			Name:     "Pakistan Rupee",
			Decimals: 100,
		},
		"590": Currency{
			Code:     "pab",
			Country:  "PANAMA",
			Name:     "Balboa",
			Decimals: 100,
		},
		"598": Currency{
			Code:     "pgk",
			Country:  "PAPUA NEW GUINEA",
			Name:     "Kina",
			Decimals: 100,
		},
		"600": Currency{
			Code:     "pyg",
			Country:  "PARAGUAY",
			Name:     "Guarani",
			Decimals: 1,
		},
		"604": Currency{
			Code:     "pen",
			Country:  "PERU",
			Name:     "Sol",
			Decimals: 100,
		},
		"608": Currency{
			Code:     "php",
			Country:  "PHILIPPINES (THE)",
			Name:     "Philippine Peso",
			Decimals: 100,
		},
		"634": Currency{
			Code:     "qar",
			Country:  "QATAR",
			Name:     "Qatari Rial",
			Decimals: 100,
		},
		"643": Currency{
			Code:     "rub",
			Country:  "RUSSIAN FEDERATION (THE)",
			Name:     "Russian Ruble",
			Decimals: 100,
		},
		"646": Currency{
			Code:     "rwf",
			Country:  "RWANDA",
			Name:     "Rwanda Franc",
			Decimals: 1,
		},
		"654": Currency{
			Code:     "shp",
			Country:  "SAINT HELENA, ASCENSION AND TRISTAN DA CUNHA",
			Name:     "Saint Helena Pound",
			Decimals: 100,
		},
		"682": Currency{
			Code:     "sar",
			Country:  "SAUDI ARABIA",
			Name:     "Saudi Riyal",
			Decimals: 100,
		},
		"690": Currency{
			Code:     "scr",
			Country:  "SEYCHELLES",
			Name:     "Seychelles Rupee",
			Decimals: 100,
		},
		"694": Currency{
			Code:     "sll",
			Country:  "SIERRA LEONE",
			Name:     "Leone",
			Decimals: 100,
		},
		"702": Currency{
			Code:     "sgd",
			Country:  "SINGAPORE",
			Name:     "Singapore Dollar",
			Decimals: 100,
		},
		"704": Currency{
			Code:     "vnd",
			Country:  "VIET NAM",
			Name:     "Dong",
			Decimals: 1,
		},
		"706": Currency{
			Code:     "sos",
			Country:  "SOMALIA",
			Name:     "Somali Shilling",
			Decimals: 100,
		},
		"710": Currency{
			Code:     "zar",
			Country:  "SOUTH AFRICA",
			Name:     "Rand",
			Decimals: 100,
		},
		"728": Currency{
			Code:     "ssp",
			Country:  "SOUTH SUDAN",
			Name:     "South Sudanese Pound",
			Decimals: 100,
		},
		"748": Currency{
			Code:     "szl",
			Country:  "ESWATINI",
			Name:     "Lilangeni",
			Decimals: 100,
		},
		"752": Currency{
			Code:     "sek",
			Country:  "SWEDEN",
			Name:     "Swedish Krona",
			Decimals: 100,
		},
		"756": Currency{
			Code:     "chf",
			Country:  "SWITZERLAND",
			Name:     "Swiss Franc",
			Decimals: 100,
		},
		"760": Currency{
			Code:     "syp",
			Country:  "SYRIAN ARAB REPUBLIC",
			Name:     "Syrian Pound",
			Decimals: 100,
		},
		"764": Currency{
			Code:     "thb",
			Country:  "THAILAND",
			Name:     "Baht",
			Decimals: 100,
		},
		"776": Currency{
			Code:     "top",
			Country:  "TONGA",
			Name:     "Pa’anga",
			Decimals: 100,
		},
		"780": Currency{
			Code:     "ttd",
			Country:  "TRINIDAD AND TOBAGO",
			Name:     "Trinidad and Tobago Dollar",
			Decimals: 100,
		},
		"784": Currency{
			Code:     "aed",
			Country:  "UNITED ARAB EMIRATES (THE)",
			Name:     "UAE Dirham",
			Decimals: 100,
		},
		"788": Currency{
			Code:     "tnd",
			Country:  "TUNISIA",
			Name:     "Tunisian Dinar",
			Decimals: 1000,
		},
		"800": Currency{
			Code:     "ugx",
			Country:  "UGANDA",
			Name:     "Uganda Shilling",
			Decimals: 1,
		},
		"807": Currency{
			Code:     "mkd",
			Country:  "MACEDONIA (THE FORMER YUGOSLAV REPUBLIC OF)",
			Name:     "Denar",
			Decimals: 100,
		},
		"818": Currency{
			Code:     "egp",
			Country:  "EGYPT",
			Name:     "Egyptian Pound",
			Decimals: 100,
		},
		"826": Currency{
			Code:     "gbp",
			Country:  "UNITED KINGDOM OF GREAT BRITAIN AND NORTHERN IRELAND (THE)",
			Name:     "Pound Sterling",
			Decimals: 100,
		},
		"834": Currency{
			Code:     "tzs",
			Country:  "TANZANIA, UNITED REPUBLIC OF",
			Name:     "Tanzanian Shilling",
			Decimals: 100,
		},
		"840": Currency{
			Code:     "usd",
			Country:  "VIRGIN ISLANDS (U.S.)",
			Name:     "US Dollar",
			Decimals: 100,
		},
		"858": Currency{
			Code:     "uyu",
			Country:  "URUGUAY",
			Name:     "Peso Uruguayo",
			Decimals: 100,
		},
		"860": Currency{
			Code:     "uzs",
			Country:  "UZBEKISTAN",
			Name:     "Uzbekistan Sum",
			Decimals: 100,
		},
		"882": Currency{
			Code:     "wst",
			Country:  "SAMOA",
			Name:     "Tala",
			Decimals: 100,
		},
		"886": Currency{
			Code:     "yer",
			Country:  "YEMEN",
			Name:     "Yemeni Rial",
			Decimals: 100,
		},
		"901": Currency{
			Code:     "twd",
			Country:  "TAIWAN (PROVINCE OF CHINA)",
			Name:     "New Taiwan Dollar",
			Decimals: 100,
		},
		"927": Currency{
			Code:     "uyw",
			Country:  "URUGUAY",
			Name:     "Unidad Previsional",
			Decimals: 10000,
		},
		"928": Currency{
			Code:     "ves",
			Country:  "VENEZUELA (BOLIVARIAN REPUBLIC OF)",
			Name:     "Bolívar Soberano",
			Decimals: 100,
		},
		"929": Currency{
			Code:     "mru",
			Country:  "MAURITANIA",
			Name:     "Ouguiya",
			Decimals: 100,
		},
		"930": Currency{
			Code:     "stn",
			Country:  "SAO TOME AND PRINCIPE",
			Name:     "Dobra",
			Decimals: 100,
		},
		"931": Currency{
			Code:     "cuc",
			Country:  "CUBA",
			Name:     "Peso Convertible",
			Decimals: 100,
		},
		"932": Currency{
			Code:     "zwl",
			Country:  "ZIMBABWE",
			Name:     "Zimbabwe Dollar",
			Decimals: 100,
		},
		"933": Currency{
			Code:     "byn",
			Country:  "BELARUS",
			Name:     "Belarusian Ruble",
			Decimals: 100,
		},
		"934": Currency{
			Code:     "tmt",
			Country:  "TURKMENISTAN",
			Name:     "Turkmenistan New Manat",
			Decimals: 100,
		},
		"936": Currency{
			Code:     "ghs",
			Country:  "GHANA",
			Name:     "Ghana Cedi",
			Decimals: 100,
		},
		"938": Currency{
			Code:     "sdg",
			Country:  "SUDAN (THE)",
			Name:     "Sudanese Pound",
			Decimals: 100,
		},
		"940": Currency{
			Code:     "uyi",
			Country:  "URUGUAY",
			Name:     "Uruguay Peso en Unidades Indexadas (UI)",
			Decimals: 1,
		},
		"941": Currency{
			Code:     "rsd",
			Country:  "SERBIA",
			Name:     "Serbian Dinar",
			Decimals: 100,
		},
		"943": Currency{
			Code:     "mzn",
			Country:  "MOZAMBIQUE",
			Name:     "Mozambique Metical",
			Decimals: 100,
		},
		"944": Currency{
			Code:     "azn",
			Country:  "AZERBAIJAN",
			Name:     "Azerbaijan Manat",
			Decimals: 100,
		},
		"946": Currency{
			Code:     "ron",
			Country:  "ROMANIA",
			Name:     "Romanian Leu",
			Decimals: 100,
		},
		"947": Currency{
			Code:     "che",
			Country:  "SWITZERLAND",
			Name:     "WIR Euro",
			Decimals: 100,
		},
		"948": Currency{
			Code:     "chw",
			Country:  "SWITZERLAND",
			Name:     "WIR Franc",
			Decimals: 100,
		},
		"949": Currency{
			Code:     "try",
			Country:  "TURKEY",
			Name:     "Turkish Lira",
			Decimals: 100,
		},
		"950": Currency{
			Code:     "xaf",
			Country:  "GABON",
			Name:     "CFA Franc BEAC",
			Decimals: 1,
		},
		"951": Currency{
			Code:     "xcd",
			Country:  "SAINT VINCENT AND THE GRENADINES",
			Name:     "East Caribbean Dollar",
			Decimals: 100,
		},
		"952": Currency{
			Code:     "xof",
			Country:  "TOGO",
			Name:     "CFA Franc BCEAO",
			Decimals: 1,
		},
		"953": Currency{
			Code:     "xpf",
			Country:  "WALLIS AND FUTUNA",
			Name:     "CFP Franc",
			Decimals: 1,
		},
		"967": Currency{
			Code:     "zmw",
			Country:  "ZAMBIA",
			Name:     "Zambian Kwacha",
			Decimals: 100,
		},
		"968": Currency{
			Code:     "srd",
			Country:  "SURINAME",
			Name:     "Surinam Dollar",
			Decimals: 100,
		},
		"969": Currency{
			Code:     "mga",
			Country:  "MADAGASCAR",
			Name:     "Malagasy Ariary",
			Decimals: 100,
		},
		"970": Currency{
			Code:     "cou",
			Country:  "COLOMBIA",
			Name:     "Unidad de Valor Real",
			Decimals: 100,
		},
		"971": Currency{
			Code:     "afn",
			Country:  "AFGHANISTAN",
			Name:     "Afghani",
			Decimals: 100,
		},
		"972": Currency{
			Code:     "tjs",
			Country:  "TAJIKISTAN",
			Name:     "Somoni",
			Decimals: 100,
		},
		"973": Currency{
			Code:     "aoa",
			Country:  "ANGOLA",
			Name:     "Kwanza",
			Decimals: 100,
		},
		"975": Currency{
			Code:     "bgn",
			Country:  "BULGARIA",
			Name:     "Bulgarian Lev",
			Decimals: 100,
		},
		"976": Currency{
			Code:     "cdf",
			Country:  "CONGO (THE DEMOCRATIC REPUBLIC OF THE)",
			Name:     "Congolese Franc",
			Decimals: 100,
		},
		"977": Currency{
			Code:     "bam",
			Country:  "BOSNIA AND HERZEGOVINA",
			Name:     "Convertible Mark",
			Decimals: 100,
		},
		"978": Currency{
			Code:     "eur",
			Country:  "SPAIN",
			Name:     "Euro",
			Decimals: 100,
		},
		"979": Currency{
			Code:     "mxv",
			Country:  "MEXICO",
			Name:     "Mexican Unidad de Inversion (UDI)",
			Decimals: 100,
		},
		"980": Currency{
			Code:     "uah",
			Country:  "UKRAINE",
			Name:     "Hryvnia",
			Decimals: 100,
		},
		"981": Currency{
			Code:     "gel",
			Country:  "GEORGIA",
			Name:     "Lari",
			Decimals: 100,
		},
		"984": Currency{
			Code:     "bov",
			Country:  "BOLIVIA (PLURINATIONAL STATE OF)",
			Name:     "Mvdol",
			Decimals: 100,
		},
		"985": Currency{
			Code:     "pln",
			Country:  "POLAND",
			Name:     "Zloty",
			Decimals: 100,
		},
		"986": Currency{
			Code:     "brl",
			Country:  "BRAZIL",
			Name:     "Brazilian Real",
			Decimals: 100,
		},
		"990": Currency{
			Code:     "clf",
			Country:  "CHILE",
			Name:     "Unidad de Fomento",
			Decimals: 10000,
		},
		"997": Currency{
			Code:     "usn",
			Country:  "UNITED STATES OF AMERICA (THE)",
			Name:     "US Dollar (Next day)",
			Decimals: 100,
		},
		"008": Currency{
			Code:     "all",
			Country:  "ALBANIA",
			Name:     "Lek",
			Decimals: 100,
		},
		"012": Currency{
			Code:     "dzd",
			Country:  "ALGERIA",
			Name:     "Algerian Dinar",
			Decimals: 100,
		},
		"032": Currency{
			Code:     "ars",
			Country:  "ARGENTINA",
			Name:     "Argentine Peso",
			Decimals: 100,
		},
		"051": Currency{
			Code:     "amd",
			Country:  "ARMENIA",
			Name:     "Armenian Dram",
			Decimals: 100,
		},
		"036": Currency{
			Code:     "aud",
			Country:  "TUVALU",
			Name:     "Australian Dollar",
			Decimals: 100,
		},
		"044": Currency{
			Code:     "bsd",
			Country:  "BAHAMAS (THE)",
			Name:     "Bahamian Dollar",
			Decimals: 100,
		},
		"048": Currency{
			Code:     "bhd",
			Country:  "BAHRAIN",
			Name:     "Bahraini Dinar",
			Decimals: 1000,
		},
		"050": Currency{
			Code:     "bdt",
			Country:  "BANGLADESH",
			Name:     "Taka",
			Decimals: 100,
		},
		"052": Currency{
			Code:     "bbd",
			Country:  "BARBADOS",
			Name:     "Barbados Dollar",
			Decimals: 100,
		},
		"084": Currency{
			Code:     "bzd",
			Country:  "BELIZE",
			Name:     "Belize Dollar",
			Decimals: 100,
		},
		"060": Currency{
			Code:     "bmd",
			Country:  "BERMUDA",
			Name:     "Bermudian Dollar",
			Decimals: 100,
		},
		"064": Currency{
			Code:     "btn",
			Country:  "BHUTAN",
			Name:     "Ngultrum",
			Decimals: 100,
		},
		"068": Currency{
			Code:     "bob",
			Country:  "BOLIVIA (PLURINATIONAL STATE OF)",
			Name:     "Boliviano",
			Decimals: 100,
		},
		"072": Currency{
			Code:     "bwp",
			Country:  "BOTSWANA",
			Name:     "Pula",
			Decimals: 100,
		},
		"096": Currency{
			Code:     "bnd",
			Country:  "BRUNEI DARUSSALAM",
			Name:     "Brunei Dollar",
			Decimals: 100,
		},
		"090": Currency{
			Code:     "sbd",
			Country:  "SOLOMON ISLANDS",
			Name:     "Solomon Islands Dollar",
			Decimals: 100,
		},
	}
)

type (
	// IsoStruct is an iso8583 container
	Message struct {
		spec   Spec
		mti    string
		bitmap []byte
		data   map[int]Field
	}

	Field struct {
		ID        string
		Value     string //the decoded value that can be used as is
		Subfields map[int]Field
	}

	// FieldDescription contains fields that describes an iso8583 Field
	FieldDescription struct {
		ContentType string                   `yaml:"ContentType"`
		MaxLen      int                      `yaml:"MaxLen"`
		MinLen      int                      `yaml:"MinLen"`
		LenType     string                   `yaml:"LenType"`
		Label       string                   `yaml:"Label"`
		Format      string                   `yaml:"Format"`
		Subfields   map[int]FieldDescription `yaml:"Subfields"`
	}

	// Spec contains a strutured description of an iso8583 spec
	// properly defined by a spec file
	Spec struct {
		version      string
		fields       map[int]FieldDescription
		messageFlows map[string]MessageFlow
	}

	MessageFlow struct {
		RepeaterMTI     string
		ResponseMTI     string
		Response        string
		Context         string
		Flow            string
		Name            string
		Source          string
		Destination     string
		Description     string
		MandatoryFields map[int]string
	}

	MandatoryFields struct {
		Fields map[string]map[int]string
	}

	// ValidationError happens when validation fails
	ValidationError struct {
		message string
	}
	Currency struct {
		Code     string
		Name     string
		Country  string
		Number   string
		Decimals int
	}
)
