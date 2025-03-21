package swiftmsgapi 
import (
	"os"
	"strings"
)

// LanguageMap maps common language names and variants to ISO language codes
var LanguageMap = map[string]string{
	// English variants
	"english":     "en",
	"en":          "en",
	"eng":         "en",
	"en_us":       "en",
	"en_gb":       "en",
	"en_uk":       "en",
	"en_ca":       "en",
	"en_au":       "en",
	
	// Spanish variants
	"spanish":     "es",
	"es":          "es",
	"esp":         "es",
	"español":     "es",
	"espanol":     "es",
	"es_es":       "es",
	"es_mx":       "es",
	"es_ar":       "es",
	
	// French variants
	"french":      "fr",
	"fr":          "fr",
	"fra":         "fr",
	"français":    "fr",
	"francais":    "fr",
	"fr_fr":       "fr",
	"fr_ca":       "fr",
	"fr_be":       "fr",
	
	// German variants
	"german":      "de",
	"de":          "de",
	"deu":         "de",
	"deutsch":     "de",
	"de_de":       "de",
	"de_at":       "de",
	"de_ch":       "de",
	
	// Italian variants
	"italian":     "it",
	"it":          "it",
	"ita":         "it",
	"italiano":    "it",
	"it_it":       "it",
	"it_ch":       "it",
	
	// Portuguese variants
	"portuguese":  "pt",
	"pt":          "pt",
	"por":         "pt",
	"português":   "pt",
	"portugues":   "pt",
	"pt_pt":       "pt",
	"pt_br":       "pt",
	
	// Chinese variants
	"chinese":     "zh",
	"zh":          "zh",
	"zho":         "zh",
	"mandarin":    "zh",
	"zh_cn":       "zh",
	"zh_tw":       "zh",
	"zh_hk":       "zh",
	
	// Japanese variants
	"japanese":    "ja",
	"ja":          "ja",
	"jpn":         "ja",
	"日本語":      "ja",
	"ja_jp":       "ja",
	
	// Korean variants
	"korean":      "ko",
	"ko":          "ko",
	"kor":         "ko",
	"한국어":      "ko",
	"ko_kr":       "ko",
	
	// Russian variants
	"russian":     "ru",
	"ru":          "ru",
	"rus":         "ru",
	"русский":     "ru",
	"ru_ru":       "ru",
	
	// Arabic variants
	"arabic":      "ar",
	"ar":          "ar",
	"ara":         "ar",
	"العربية":     "ar",
	"ar_sa":       "ar",
	
	// Hindi variants
	"hindi":       "hi",
	"hi":          "hi",
	"hin":         "hi",
	"हिन्दी":       "hi",
	"hi_in":       "hi",
	
	// Other common languages
	"dutch":       "nl",
	"nl":          "nl",
	"nederlands":  "nl",
	
	"swedish":     "sv",
	"sv":          "sv",
	"svenska":     "sv",
	
	"norwegian":   "no",
	"no":          "no",
	"norsk":       "no",
	
	"finnish":     "fi",
	"fi":          "fi",
	"suomi":       "fi",
	
	"danish":      "da",
	"da":          "da",
	"dansk":       "da",
	
	"greek":       "el",
	"el":          "el",
	"ελληνικά":    "el",
	
	"turkish":     "tr",
	"tr":          "tr",
	"türkçe":      "tr",
	"turkce":      "tr",
	
	"thai":        "th",
	"th":          "th",
	"ไทย":         "th",
	
	"vietnamese":  "vi",
	"vi":          "vi",
	"tiếng việt":  "vi",
	"tieng viet":  "vi",
}

// GetLanguageCode converts a language name or variant to the standard ISO code
func GetLanguageCode(input string) string {
	// Default to English if empty
	if input == "" {
		return "en"
	}
	
	// Convert to lowercase for case-insensitive matching
	normalized := strings.ToLower(strings.TrimSpace(input))
	
	// Check if it's in our map
	if code, exists := LanguageMap[normalized]; exists {
		return code
	}
	
	// If input contains underscore or dot (like en_US.UTF-8)
	if strings.Contains(normalized, "_") {
		parts := strings.Split(normalized, "_")
		if len(parts) > 0 {
			// Try the first part (language code)
			if code, exists := LanguageMap[parts[0]]; exists {
				return code
			}
		}
	}
	
	// If not found in map and not parseable, return as-is
	// If it's a valid ISO code, it will work; if not, it will fail gracefully
	return normalized
}

// GetSystemLanguage determines the system language from environment variables
func GetSystemLanguage() string {
	// Check environment variables in order of precedence
	for _, envVar := range []string{"LC_ALL", "LC_MESSAGES", "LANG", "LANGUAGE"} {
		locale := os.Getenv(envVar)
		if locale != "" {
			// Extract language code (e.g., "en" from "en_US.UTF-8")
			if strings.Contains(locale, "_") {
				return strings.Split(locale, "_")[0]
			}
			if strings.Contains(locale, ".") {
				return strings.Split(locale, ".")[0]
			}
			return locale
		}
	}
	
	// Default to English if no locale is set
	return "en"
}
