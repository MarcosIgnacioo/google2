package parserfake

func str_fn(arg string) string       {}
func int_fn(arg int) int             {}
func float32_fn(arg float32) float32 {}
func float64_fn(arg float64) float64 {}
func bool_fn(arg bool) bool          {}
func byte_fn(arg byte) byte          {}
func rune_fn(arg rune) rune          {}

func str_fn_NO_RETURN(arg string)      {}
func int_fn_NO_RETURN(arg int)         {}
func float32_fn_NO_RETURN(arg float32) {}
func float64_fn_NO_RETURN(arg float64) {}
func bool_fn_NO_RETURN(arg bool)       {}
func byte_fn_NO_RETURN(arg byte)       {}
func rune_fn_NO_RETURN(arg rune)       {}

func arr_str_fn(arg []string) []string       {}
func arr_int_fn(arg []int) []int             {}
func arr_float32_fn(arg []float32) []float32 {}
func arr_float64_fn(arg []float64) []float64 {}
func arr_bool_fn(arg []bool) []bool          {}
func arr_byte_fn(arg []byte) []byte          {}
func arr_rune_fn(arg []rune) []rune          {}

func arr_str_fn_NO_RETURN(arg []string)      {}
func arr_int_fn_NO_RETURN(arg []int)         {}
func arr_float32_fn_NO_RETURN(arg []float32) {}
func arr_float64_fn_NO_RETURN(arg []float64) {}
func arr_bool_fn_NO_RETURN(arg []bool)       {}
func arr_byte_fn_NO_RETURN(arg []byte)       {}
func arr_rune_fn_NO_RETURN(arg []rune)       {}

func (X *string) method_str_fn(arg string) string        {}
func (X *int) method_int_fn(arg int) int                 {}
func (X *float32) method_float32_fn(arg float32) float32 {}
func (X *float64) method_float64_fn(arg float64) float64 {}
func (X *bool) method_bool_fn(arg bool) bool             {}
func (X *byte) method_byte_fn(arg byte) byte             {}
func (X *rune) method_rune_fn(arg rune) rune             {}

func (X *string) method_str_fn_NO_RETURN(arg string)       {}
func (X *int) method_int_fn_NO_RETURN(arg int)             {}
func (X *float32) method_float32_fn_NO_RETURN(arg float32) {}
func (X *float64) method_float64_fn_NO_RETURN(arg float64) {}
func (X *bool) method_bool_fn_NO_RETURN(arg bool)          {}
func (X *byte) method_byte_fn_NO_RETURN(arg byte)          {}
func (X *rune) method_rune_fn_NO_RETURN(arg rune)          {}

func generics_str_fn[K string](arg string) string        {}
func generics_int_fn[K int](arg int) int                 {}
func generics_float32_fn[K float32](arg float32) float32 {}
func generics_float64_fn[K float64](arg float64) float64 {}
func generics_bool_fn[K bool](arg bool) bool             {}
func generics_byte_fn[K byte](arg byte) byte             {}
func generics_rune_fn[K rune](arg rune) rune             {}

func generics_str_fn_NO_RETURN[K string](arg string)       {}
func generics_int_fn_NO_RETURN[K int](arg int)             {}
func generics_float32_fn_NO_RETURN[K float32](arg float32) {}
func generics_float64_fn_NO_RETURN[K float64](arg float64) {}
func generics_bool_fn_NO_RETURN[K bool](arg bool)          {}
func generics_byte_fn_NO_RETURN[K byte](arg byte)          {}
func generics_rune_fn_NO_RETURN[K rune](arg rune)          {}

func generics_2_str_fn[K string | Foo](arg string) string        {}
func generics_2_int_fn[K int | Foo](arg int) int                 {}
func generics_2_float32_fn[K float32 | Foo](arg float32) float32 {}
func generics_2_float64_fn[K float64 | Foo](arg float64) float64 {}
func generics_2_bool_fn[K bool | Foo](arg bool) bool             {}
func generics_2_byte_fn[K byte | Foo](arg byte) byte             {}
func generics_2_rune_fn[K rune | Foo](arg rune) rune             {}

func generics_2_str_fn_NO_RETURN[K string | Foo](arg string)       {}
func generics_2_int_fn_NO_RETURN[K int | Foo](arg int)             {}
func generics_2_float32_fn_NO_RETURN[K float32 | Foo](arg float32) {}
func generics_2_float64_fn_NO_RETURN[K float64 | Foo](arg float64) {}
func generics_2_bool_fn_NO_RETURN[K bool | Foo](arg bool)          {}
func generics_2_byte_fn_NO_RETURN[K byte | Foo](arg byte)          {}
func generics_2_rune_fn_NO_RETURN[K rune | Foo](arg rune)          {}
