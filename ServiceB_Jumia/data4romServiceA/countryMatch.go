package data4romServiceA

func matchCodeToCountry(code int) string{
    var countryCode = map[string]int{
         254         : "KENYA",
         237         : "CAMEROON",
         251         : "ERITREA",
         212         : "MOZAMBIQUE",
         256         : "UGANDA",
         233         : "GHANA",
         234         : "NIGERIA",
         27          : "SOUTHAFRICA",
         216         : "TUNISIA",
         251         : "ETHIOPIA",
         213         : "RWANDA",
         230         : "MAURITIUS",
         253         : "DJIBOUTI",
         244         : "ANGOLA",
         236         : "CENTRALAFRICANREPUBLIC",
         235         : "CHAD",
         242         : "CONGO",
         240         : "EQUATORIALGUINEA",
         241         :  "GABON",
         239         : "SAOTOME_REPUBLIC",
         269         : "COMOROS",
         265         : "MALAWI",
         250         : "RWANDA",
         252         : "SOMALIA",
         258         : "MOZAMBIQUE",
         248         : "SEYCHELLES",
         220         : "GAMBIA",
         224         : "GUINEA",
         223         : "MALI",
         221         : "SENEGAL",
         227           "NIGER",
         231         : "LIBERIA",
         260         :  "ZAMBIA",
         228         :  "TOGO",
         225         : "IVORY_COST",
         238         : "CAPE VERDE",
         229         : "BENIN",
         245         :  "GUINEA_BISAU",
         222         :  "MAURITANIA",
         232         : "SIERA_LEON",
         226         : "BUKINA_FASO",
         249         : "SUDAN",
         20          : "EGYPT",
         218         : "LIBYA"
      }

      return countryCode[int]

}
