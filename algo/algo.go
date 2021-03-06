//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

package algo ;import _b "strconv";func _e (_g byte )bool {return _g >='0'&&_g <='9'};func RepeatString (s string ,cnt int )string {if cnt <=0{return "";};_dd :=make ([]byte ,len (s )*cnt );_gd :=[]byte (s );for _f :=0;_f < cnt ;_f ++{copy (_dd [_f :],_gd );};return string (_dd );};

// NaturalLess compares two strings in a human manner so rId2 sorts less than rId10
func NaturalLess (lhs ,rhs string )bool {_a ,_bf :=0,0;for _a < len (lhs )&&_bf < len (rhs ){_c :=lhs [_a ];_bd :=rhs [_bf ];_df :=_e (_c );_ag :=_e (_bd );switch {case _df &&!_ag :return true ;case !_df &&_ag :return false ;case !_df &&!_ag :if _c !=_bd {return _c < _bd ;};_a ++;_bf ++;default:_ea :=_a +1;_af :=_bf +1;for _ea < len (lhs )&&_e (lhs [_ea ]){_ea ++;};for _af < len (rhs )&&_e (rhs [_af ]){_af ++;};_gc ,_ :=_b .ParseUint (lhs [_a :_ea ],10,64);_da ,_ :=_b .ParseUint (rhs [_a :_af ],10,64);if _gc !=_da {return _gc < _da ;};_a =_ea ;_bf =_af ;};};return len (lhs )< len (rhs );};