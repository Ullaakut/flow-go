// Code generated from parser/Strictus.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser

import (
	"fmt"
	"unicode"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

// Suppress unused import error
var _ = fmt.Printf
var _ = unicode.IsLetter

var serializedLexerAtn = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 2, 62, 447,
	8, 1, 4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7,
	9, 7, 4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12,
	4, 13, 9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4,
	18, 9, 18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23,
	9, 23, 4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9,
	28, 4, 29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33,
	4, 34, 9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 4,
	39, 9, 39, 4, 40, 9, 40, 4, 41, 9, 41, 4, 42, 9, 42, 4, 43, 9, 43, 4, 44,
	9, 44, 4, 45, 9, 45, 4, 46, 9, 46, 4, 47, 9, 47, 4, 48, 9, 48, 4, 49, 9,
	49, 4, 50, 9, 50, 4, 51, 9, 51, 4, 52, 9, 52, 4, 53, 9, 53, 4, 54, 9, 54,
	4, 55, 9, 55, 4, 56, 9, 56, 4, 57, 9, 57, 4, 58, 9, 58, 4, 59, 9, 59, 4,
	60, 9, 60, 4, 61, 9, 61, 4, 62, 9, 62, 4, 63, 9, 63, 4, 64, 9, 64, 4, 65,
	9, 65, 4, 66, 9, 66, 3, 2, 3, 2, 3, 3, 3, 3, 3, 4, 3, 4, 3, 5, 3, 5, 3,
	6, 3, 6, 3, 7, 3, 7, 3, 8, 3, 8, 3, 9, 3, 9, 3, 9, 3, 10, 3, 10, 3, 10,
	3, 11, 3, 11, 3, 12, 3, 12, 3, 13, 3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 17, 3, 17, 3, 17, 3, 18, 3, 18, 3, 18, 3, 19,
	3, 19, 3, 20, 3, 20, 3, 21, 3, 21, 3, 22, 3, 22, 3, 23, 3, 23, 3, 24, 3,
	24, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28,
	3, 28, 3, 29, 3, 29, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3,
	30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31,
	3, 31, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3, 32, 3,
	32, 3, 33, 3, 33, 3, 33, 3, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37, 3,
	37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 39, 3, 40, 3, 40, 3,
	40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 40, 3, 41, 3, 41, 3, 41, 3, 41,
	3, 42, 3, 42, 3, 42, 3, 42, 3, 43, 3, 43, 3, 43, 3, 44, 3, 44, 3, 44, 3,
	44, 3, 44, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 45, 3, 46, 3, 46, 3, 46,
	3, 46, 3, 46, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 47, 3, 48, 3, 48, 3,
	48, 3, 48, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 49, 3, 50, 3, 50,
	3, 50, 3, 50, 3, 50, 3, 51, 3, 51, 7, 51, 327, 10, 51, 12, 51, 14, 51,
	330, 11, 51, 3, 52, 5, 52, 333, 10, 52, 3, 53, 3, 53, 5, 53, 337, 10, 53,
	3, 54, 3, 54, 7, 54, 341, 10, 54, 12, 54, 14, 54, 344, 11, 54, 3, 55, 3,
	55, 3, 55, 3, 55, 6, 55, 350, 10, 55, 13, 55, 14, 55, 351, 3, 56, 3, 56,
	3, 56, 3, 56, 6, 56, 358, 10, 56, 13, 56, 14, 56, 359, 3, 57, 3, 57, 3,
	57, 3, 57, 6, 57, 366, 10, 57, 13, 57, 14, 57, 367, 3, 58, 3, 58, 3, 58,
	7, 58, 373, 10, 58, 12, 58, 14, 58, 376, 11, 58, 3, 59, 3, 59, 7, 59, 380,
	10, 59, 12, 59, 14, 59, 383, 11, 59, 3, 59, 3, 59, 3, 60, 3, 60, 5, 60,
	389, 10, 60, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 3, 61, 6, 61, 398,
	10, 61, 13, 61, 14, 61, 399, 3, 61, 3, 61, 5, 61, 404, 10, 61, 3, 62, 3,
	62, 3, 63, 6, 63, 409, 10, 63, 13, 63, 14, 63, 410, 3, 63, 3, 63, 3, 64,
	6, 64, 416, 10, 64, 13, 64, 14, 64, 417, 3, 64, 3, 64, 3, 65, 3, 65, 3,
	65, 3, 65, 3, 65, 7, 65, 427, 10, 65, 12, 65, 14, 65, 430, 11, 65, 3, 65,
	3, 65, 3, 65, 3, 65, 3, 65, 3, 66, 3, 66, 3, 66, 3, 66, 7, 66, 441, 10,
	66, 12, 66, 14, 66, 444, 11, 66, 3, 66, 3, 66, 3, 428, 2, 67, 3, 3, 5,
	4, 7, 5, 9, 6, 11, 7, 13, 8, 15, 9, 17, 10, 19, 11, 21, 12, 23, 13, 25,
	14, 27, 15, 29, 16, 31, 17, 33, 18, 35, 19, 37, 20, 39, 21, 41, 22, 43,
	23, 45, 24, 47, 25, 49, 26, 51, 27, 53, 28, 55, 29, 57, 30, 59, 31, 61,
	32, 63, 33, 65, 34, 67, 35, 69, 36, 71, 37, 73, 38, 75, 39, 77, 40, 79,
	41, 81, 42, 83, 43, 85, 44, 87, 45, 89, 46, 91, 47, 93, 48, 95, 49, 97,
	50, 99, 51, 101, 52, 103, 2, 105, 2, 107, 53, 109, 54, 111, 55, 113, 56,
	115, 57, 117, 58, 119, 2, 121, 2, 123, 2, 125, 59, 127, 60, 129, 61, 131,
	62, 3, 2, 15, 5, 2, 67, 92, 97, 97, 99, 124, 3, 2, 50, 59, 4, 2, 50, 59,
	97, 97, 4, 2, 50, 51, 97, 97, 4, 2, 50, 57, 97, 97, 6, 2, 50, 59, 67, 72,
	97, 97, 99, 104, 4, 2, 67, 92, 99, 124, 6, 2, 50, 59, 67, 92, 97, 97, 99,
	124, 6, 2, 12, 12, 15, 15, 36, 36, 94, 94, 9, 2, 36, 36, 41, 41, 50, 50,
	94, 94, 112, 112, 116, 116, 118, 118, 5, 2, 50, 59, 67, 72, 99, 104, 6,
	2, 2, 2, 11, 11, 13, 14, 34, 34, 4, 2, 12, 12, 15, 15, 2, 457, 2, 3, 3,
	2, 2, 2, 2, 5, 3, 2, 2, 2, 2, 7, 3, 2, 2, 2, 2, 9, 3, 2, 2, 2, 2, 11, 3,
	2, 2, 2, 2, 13, 3, 2, 2, 2, 2, 15, 3, 2, 2, 2, 2, 17, 3, 2, 2, 2, 2, 19,
	3, 2, 2, 2, 2, 21, 3, 2, 2, 2, 2, 23, 3, 2, 2, 2, 2, 25, 3, 2, 2, 2, 2,
	27, 3, 2, 2, 2, 2, 29, 3, 2, 2, 2, 2, 31, 3, 2, 2, 2, 2, 33, 3, 2, 2, 2,
	2, 35, 3, 2, 2, 2, 2, 37, 3, 2, 2, 2, 2, 39, 3, 2, 2, 2, 2, 41, 3, 2, 2,
	2, 2, 43, 3, 2, 2, 2, 2, 45, 3, 2, 2, 2, 2, 47, 3, 2, 2, 2, 2, 49, 3, 2,
	2, 2, 2, 51, 3, 2, 2, 2, 2, 53, 3, 2, 2, 2, 2, 55, 3, 2, 2, 2, 2, 57, 3,
	2, 2, 2, 2, 59, 3, 2, 2, 2, 2, 61, 3, 2, 2, 2, 2, 63, 3, 2, 2, 2, 2, 65,
	3, 2, 2, 2, 2, 67, 3, 2, 2, 2, 2, 69, 3, 2, 2, 2, 2, 71, 3, 2, 2, 2, 2,
	73, 3, 2, 2, 2, 2, 75, 3, 2, 2, 2, 2, 77, 3, 2, 2, 2, 2, 79, 3, 2, 2, 2,
	2, 81, 3, 2, 2, 2, 2, 83, 3, 2, 2, 2, 2, 85, 3, 2, 2, 2, 2, 87, 3, 2, 2,
	2, 2, 89, 3, 2, 2, 2, 2, 91, 3, 2, 2, 2, 2, 93, 3, 2, 2, 2, 2, 95, 3, 2,
	2, 2, 2, 97, 3, 2, 2, 2, 2, 99, 3, 2, 2, 2, 2, 101, 3, 2, 2, 2, 2, 107,
	3, 2, 2, 2, 2, 109, 3, 2, 2, 2, 2, 111, 3, 2, 2, 2, 2, 113, 3, 2, 2, 2,
	2, 115, 3, 2, 2, 2, 2, 117, 3, 2, 2, 2, 2, 125, 3, 2, 2, 2, 2, 127, 3,
	2, 2, 2, 2, 129, 3, 2, 2, 2, 2, 131, 3, 2, 2, 2, 3, 133, 3, 2, 2, 2, 5,
	135, 3, 2, 2, 2, 7, 137, 3, 2, 2, 2, 9, 139, 3, 2, 2, 2, 11, 141, 3, 2,
	2, 2, 13, 143, 3, 2, 2, 2, 15, 145, 3, 2, 2, 2, 17, 147, 3, 2, 2, 2, 19,
	150, 3, 2, 2, 2, 21, 153, 3, 2, 2, 2, 23, 155, 3, 2, 2, 2, 25, 157, 3,
	2, 2, 2, 27, 160, 3, 2, 2, 2, 29, 163, 3, 2, 2, 2, 31, 165, 3, 2, 2, 2,
	33, 167, 3, 2, 2, 2, 35, 170, 3, 2, 2, 2, 37, 173, 3, 2, 2, 2, 39, 175,
	3, 2, 2, 2, 41, 177, 3, 2, 2, 2, 43, 179, 3, 2, 2, 2, 45, 181, 3, 2, 2,
	2, 47, 183, 3, 2, 2, 2, 49, 185, 3, 2, 2, 2, 51, 187, 3, 2, 2, 2, 53, 190,
	3, 2, 2, 2, 55, 194, 3, 2, 2, 2, 57, 196, 3, 2, 2, 2, 59, 198, 3, 2, 2,
	2, 61, 210, 3, 2, 2, 2, 63, 217, 3, 2, 2, 2, 65, 227, 3, 2, 2, 2, 67, 231,
	3, 2, 2, 2, 69, 235, 3, 2, 2, 2, 71, 240, 3, 2, 2, 2, 73, 244, 3, 2, 2,
	2, 75, 253, 3, 2, 2, 2, 77, 260, 3, 2, 2, 2, 79, 266, 3, 2, 2, 2, 81, 275,
	3, 2, 2, 2, 83, 279, 3, 2, 2, 2, 85, 283, 3, 2, 2, 2, 87, 286, 3, 2, 2,
	2, 89, 291, 3, 2, 2, 2, 91, 297, 3, 2, 2, 2, 93, 302, 3, 2, 2, 2, 95, 308,
	3, 2, 2, 2, 97, 312, 3, 2, 2, 2, 99, 319, 3, 2, 2, 2, 101, 324, 3, 2, 2,
	2, 103, 332, 3, 2, 2, 2, 105, 336, 3, 2, 2, 2, 107, 338, 3, 2, 2, 2, 109,
	345, 3, 2, 2, 2, 111, 353, 3, 2, 2, 2, 113, 361, 3, 2, 2, 2, 115, 369,
	3, 2, 2, 2, 117, 377, 3, 2, 2, 2, 119, 388, 3, 2, 2, 2, 121, 403, 3, 2,
	2, 2, 123, 405, 3, 2, 2, 2, 125, 408, 3, 2, 2, 2, 127, 415, 3, 2, 2, 2,
	129, 421, 3, 2, 2, 2, 131, 436, 3, 2, 2, 2, 133, 134, 7, 46, 2, 2, 134,
	4, 3, 2, 2, 2, 135, 136, 7, 125, 2, 2, 136, 6, 3, 2, 2, 2, 137, 138, 7,
	127, 2, 2, 138, 8, 3, 2, 2, 2, 139, 140, 7, 60, 2, 2, 140, 10, 3, 2, 2,
	2, 141, 142, 7, 93, 2, 2, 142, 12, 3, 2, 2, 2, 143, 144, 7, 95, 2, 2, 144,
	14, 3, 2, 2, 2, 145, 146, 7, 63, 2, 2, 146, 16, 3, 2, 2, 2, 147, 148, 7,
	126, 2, 2, 148, 149, 7, 126, 2, 2, 149, 18, 3, 2, 2, 2, 150, 151, 7, 40,
	2, 2, 151, 152, 7, 40, 2, 2, 152, 20, 3, 2, 2, 2, 153, 154, 7, 48, 2, 2,
	154, 22, 3, 2, 2, 2, 155, 156, 7, 61, 2, 2, 156, 24, 3, 2, 2, 2, 157, 158,
	7, 63, 2, 2, 158, 159, 7, 63, 2, 2, 159, 26, 3, 2, 2, 2, 160, 161, 7, 35,
	2, 2, 161, 162, 7, 63, 2, 2, 162, 28, 3, 2, 2, 2, 163, 164, 7, 62, 2, 2,
	164, 30, 3, 2, 2, 2, 165, 166, 7, 64, 2, 2, 166, 32, 3, 2, 2, 2, 167, 168,
	7, 62, 2, 2, 168, 169, 7, 63, 2, 2, 169, 34, 3, 2, 2, 2, 170, 171, 7, 64,
	2, 2, 171, 172, 7, 63, 2, 2, 172, 36, 3, 2, 2, 2, 173, 174, 7, 45, 2, 2,
	174, 38, 3, 2, 2, 2, 175, 176, 7, 47, 2, 2, 176, 40, 3, 2, 2, 2, 177, 178,
	7, 44, 2, 2, 178, 42, 3, 2, 2, 2, 179, 180, 7, 49, 2, 2, 180, 44, 3, 2,
	2, 2, 181, 182, 7, 39, 2, 2, 182, 46, 3, 2, 2, 2, 183, 184, 7, 35, 2, 2,
	184, 48, 3, 2, 2, 2, 185, 186, 7, 65, 2, 2, 186, 50, 3, 2, 2, 2, 187, 188,
	7, 65, 2, 2, 188, 189, 7, 65, 2, 2, 189, 52, 3, 2, 2, 2, 190, 191, 7, 99,
	2, 2, 191, 192, 7, 117, 2, 2, 192, 193, 7, 65, 2, 2, 193, 54, 3, 2, 2,
	2, 194, 195, 7, 42, 2, 2, 195, 56, 3, 2, 2, 2, 196, 197, 7, 43, 2, 2, 197,
	58, 3, 2, 2, 2, 198, 199, 7, 118, 2, 2, 199, 200, 7, 116, 2, 2, 200, 201,
	7, 99, 2, 2, 201, 202, 7, 112, 2, 2, 202, 203, 7, 117, 2, 2, 203, 204,
	7, 99, 2, 2, 204, 205, 7, 101, 2, 2, 205, 206, 7, 118, 2, 2, 206, 207,
	7, 107, 2, 2, 207, 208, 7, 113, 2, 2, 208, 209, 7, 112, 2, 2, 209, 60,
	3, 2, 2, 2, 210, 211, 7, 117, 2, 2, 211, 212, 7, 118, 2, 2, 212, 213, 7,
	116, 2, 2, 213, 214, 7, 119, 2, 2, 214, 215, 7, 101, 2, 2, 215, 216, 7,
	118, 2, 2, 216, 62, 3, 2, 2, 2, 217, 218, 7, 107, 2, 2, 218, 219, 7, 112,
	2, 2, 219, 220, 7, 118, 2, 2, 220, 221, 7, 103, 2, 2, 221, 222, 7, 116,
	2, 2, 222, 223, 7, 104, 2, 2, 223, 224, 7, 99, 2, 2, 224, 225, 7, 101,
	2, 2, 225, 226, 7, 103, 2, 2, 226, 64, 3, 2, 2, 2, 227, 228, 7, 104, 2,
	2, 228, 229, 7, 119, 2, 2, 229, 230, 7, 112, 2, 2, 230, 66, 3, 2, 2, 2,
	231, 232, 7, 114, 2, 2, 232, 233, 7, 116, 2, 2, 233, 234, 7, 103, 2, 2,
	234, 68, 3, 2, 2, 2, 235, 236, 7, 114, 2, 2, 236, 237, 7, 113, 2, 2, 237,
	238, 7, 117, 2, 2, 238, 239, 7, 118, 2, 2, 239, 70, 3, 2, 2, 2, 240, 241,
	7, 114, 2, 2, 241, 242, 7, 119, 2, 2, 242, 243, 7, 100, 2, 2, 243, 72,
	3, 2, 2, 2, 244, 245, 7, 114, 2, 2, 245, 246, 7, 119, 2, 2, 246, 247, 7,
	100, 2, 2, 247, 248, 7, 42, 2, 2, 248, 249, 7, 117, 2, 2, 249, 250, 7,
	103, 2, 2, 250, 251, 7, 118, 2, 2, 251, 252, 7, 43, 2, 2, 252, 74, 3, 2,
	2, 2, 253, 254, 7, 116, 2, 2, 254, 255, 7, 103, 2, 2, 255, 256, 7, 118,
	2, 2, 256, 257, 7, 119, 2, 2, 257, 258, 7, 116, 2, 2, 258, 259, 7, 112,
	2, 2, 259, 76, 3, 2, 2, 2, 260, 261, 7, 100, 2, 2, 261, 262, 7, 116, 2,
	2, 262, 263, 7, 103, 2, 2, 263, 264, 7, 99, 2, 2, 264, 265, 7, 109, 2,
	2, 265, 78, 3, 2, 2, 2, 266, 267, 7, 101, 2, 2, 267, 268, 7, 113, 2, 2,
	268, 269, 7, 112, 2, 2, 269, 270, 7, 118, 2, 2, 270, 271, 7, 107, 2, 2,
	271, 272, 7, 112, 2, 2, 272, 273, 7, 119, 2, 2, 273, 274, 7, 103, 2, 2,
	274, 80, 3, 2, 2, 2, 275, 276, 7, 110, 2, 2, 276, 277, 7, 103, 2, 2, 277,
	278, 7, 118, 2, 2, 278, 82, 3, 2, 2, 2, 279, 280, 7, 120, 2, 2, 280, 281,
	7, 99, 2, 2, 281, 282, 7, 116, 2, 2, 282, 84, 3, 2, 2, 2, 283, 284, 7,
	107, 2, 2, 284, 285, 7, 104, 2, 2, 285, 86, 3, 2, 2, 2, 286, 287, 7, 103,
	2, 2, 287, 288, 7, 110, 2, 2, 288, 289, 7, 117, 2, 2, 289, 290, 7, 103,
	2, 2, 290, 88, 3, 2, 2, 2, 291, 292, 7, 121, 2, 2, 292, 293, 7, 106, 2,
	2, 293, 294, 7, 107, 2, 2, 294, 295, 7, 110, 2, 2, 295, 296, 7, 103, 2,
	2, 296, 90, 3, 2, 2, 2, 297, 298, 7, 118, 2, 2, 298, 299, 7, 116, 2, 2,
	299, 300, 7, 119, 2, 2, 300, 301, 7, 103, 2, 2, 301, 92, 3, 2, 2, 2, 302,
	303, 7, 104, 2, 2, 303, 304, 7, 99, 2, 2, 304, 305, 7, 110, 2, 2, 305,
	306, 7, 117, 2, 2, 306, 307, 7, 103, 2, 2, 307, 94, 3, 2, 2, 2, 308, 309,
	7, 112, 2, 2, 309, 310, 7, 107, 2, 2, 310, 311, 7, 110, 2, 2, 311, 96,
	3, 2, 2, 2, 312, 313, 7, 107, 2, 2, 313, 314, 7, 111, 2, 2, 314, 315, 7,
	114, 2, 2, 315, 316, 7, 113, 2, 2, 316, 317, 7, 116, 2, 2, 317, 318, 7,
	118, 2, 2, 318, 98, 3, 2, 2, 2, 319, 320, 7, 104, 2, 2, 320, 321, 7, 116,
	2, 2, 321, 322, 7, 113, 2, 2, 322, 323, 7, 111, 2, 2, 323, 100, 3, 2, 2,
	2, 324, 328, 5, 103, 52, 2, 325, 327, 5, 105, 53, 2, 326, 325, 3, 2, 2,
	2, 327, 330, 3, 2, 2, 2, 328, 326, 3, 2, 2, 2, 328, 329, 3, 2, 2, 2, 329,
	102, 3, 2, 2, 2, 330, 328, 3, 2, 2, 2, 331, 333, 9, 2, 2, 2, 332, 331,
	3, 2, 2, 2, 333, 104, 3, 2, 2, 2, 334, 337, 9, 3, 2, 2, 335, 337, 5, 103,
	52, 2, 336, 334, 3, 2, 2, 2, 336, 335, 3, 2, 2, 2, 337, 106, 3, 2, 2, 2,
	338, 342, 9, 3, 2, 2, 339, 341, 9, 4, 2, 2, 340, 339, 3, 2, 2, 2, 341,
	344, 3, 2, 2, 2, 342, 340, 3, 2, 2, 2, 342, 343, 3, 2, 2, 2, 343, 108,
	3, 2, 2, 2, 344, 342, 3, 2, 2, 2, 345, 346, 7, 50, 2, 2, 346, 347, 7, 100,
	2, 2, 347, 349, 3, 2, 2, 2, 348, 350, 9, 5, 2, 2, 349, 348, 3, 2, 2, 2,
	350, 351, 3, 2, 2, 2, 351, 349, 3, 2, 2, 2, 351, 352, 3, 2, 2, 2, 352,
	110, 3, 2, 2, 2, 353, 354, 7, 50, 2, 2, 354, 355, 7, 113, 2, 2, 355, 357,
	3, 2, 2, 2, 356, 358, 9, 6, 2, 2, 357, 356, 3, 2, 2, 2, 358, 359, 3, 2,
	2, 2, 359, 357, 3, 2, 2, 2, 359, 360, 3, 2, 2, 2, 360, 112, 3, 2, 2, 2,
	361, 362, 7, 50, 2, 2, 362, 363, 7, 122, 2, 2, 363, 365, 3, 2, 2, 2, 364,
	366, 9, 7, 2, 2, 365, 364, 3, 2, 2, 2, 366, 367, 3, 2, 2, 2, 367, 365,
	3, 2, 2, 2, 367, 368, 3, 2, 2, 2, 368, 114, 3, 2, 2, 2, 369, 370, 7, 50,
	2, 2, 370, 374, 9, 8, 2, 2, 371, 373, 9, 9, 2, 2, 372, 371, 3, 2, 2, 2,
	373, 376, 3, 2, 2, 2, 374, 372, 3, 2, 2, 2, 374, 375, 3, 2, 2, 2, 375,
	116, 3, 2, 2, 2, 376, 374, 3, 2, 2, 2, 377, 381, 7, 36, 2, 2, 378, 380,
	5, 119, 60, 2, 379, 378, 3, 2, 2, 2, 380, 383, 3, 2, 2, 2, 381, 379, 3,
	2, 2, 2, 381, 382, 3, 2, 2, 2, 382, 384, 3, 2, 2, 2, 383, 381, 3, 2, 2,
	2, 384, 385, 7, 36, 2, 2, 385, 118, 3, 2, 2, 2, 386, 389, 5, 121, 61, 2,
	387, 389, 10, 10, 2, 2, 388, 386, 3, 2, 2, 2, 388, 387, 3, 2, 2, 2, 389,
	120, 3, 2, 2, 2, 390, 391, 7, 94, 2, 2, 391, 404, 9, 11, 2, 2, 392, 393,
	7, 94, 2, 2, 393, 394, 7, 119, 2, 2, 394, 395, 3, 2, 2, 2, 395, 397, 7,
	125, 2, 2, 396, 398, 5, 123, 62, 2, 397, 396, 3, 2, 2, 2, 398, 399, 3,
	2, 2, 2, 399, 397, 3, 2, 2, 2, 399, 400, 3, 2, 2, 2, 400, 401, 3, 2, 2,
	2, 401, 402, 7, 127, 2, 2, 402, 404, 3, 2, 2, 2, 403, 390, 3, 2, 2, 2,
	403, 392, 3, 2, 2, 2, 404, 122, 3, 2, 2, 2, 405, 406, 9, 12, 2, 2, 406,
	124, 3, 2, 2, 2, 407, 409, 9, 13, 2, 2, 408, 407, 3, 2, 2, 2, 409, 410,
	3, 2, 2, 2, 410, 408, 3, 2, 2, 2, 410, 411, 3, 2, 2, 2, 411, 412, 3, 2,
	2, 2, 412, 413, 8, 63, 2, 2, 413, 126, 3, 2, 2, 2, 414, 416, 9, 14, 2,
	2, 415, 414, 3, 2, 2, 2, 416, 417, 3, 2, 2, 2, 417, 415, 3, 2, 2, 2, 417,
	418, 3, 2, 2, 2, 418, 419, 3, 2, 2, 2, 419, 420, 8, 64, 2, 2, 420, 128,
	3, 2, 2, 2, 421, 422, 7, 49, 2, 2, 422, 423, 7, 44, 2, 2, 423, 428, 3,
	2, 2, 2, 424, 427, 5, 129, 65, 2, 425, 427, 11, 2, 2, 2, 426, 424, 3, 2,
	2, 2, 426, 425, 3, 2, 2, 2, 427, 430, 3, 2, 2, 2, 428, 429, 3, 2, 2, 2,
	428, 426, 3, 2, 2, 2, 429, 431, 3, 2, 2, 2, 430, 428, 3, 2, 2, 2, 431,
	432, 7, 44, 2, 2, 432, 433, 7, 49, 2, 2, 433, 434, 3, 2, 2, 2, 434, 435,
	8, 65, 2, 2, 435, 130, 3, 2, 2, 2, 436, 437, 7, 49, 2, 2, 437, 438, 7,
	49, 2, 2, 438, 442, 3, 2, 2, 2, 439, 441, 10, 14, 2, 2, 440, 439, 3, 2,
	2, 2, 441, 444, 3, 2, 2, 2, 442, 440, 3, 2, 2, 2, 442, 443, 3, 2, 2, 2,
	443, 445, 3, 2, 2, 2, 444, 442, 3, 2, 2, 2, 445, 446, 8, 66, 2, 2, 446,
	132, 3, 2, 2, 2, 20, 2, 328, 332, 336, 342, 351, 359, 367, 374, 381, 388,
	399, 403, 410, 417, 426, 428, 442, 3, 2, 3, 2,
}

var lexerDeserializer = antlr.NewATNDeserializer(nil)
var lexerAtn = lexerDeserializer.DeserializeFromUInt16(serializedLexerAtn)

var lexerChannelNames = []string{
	"DEFAULT_TOKEN_CHANNEL", "HIDDEN",
}

var lexerModeNames = []string{
	"DEFAULT_MODE",
}

var lexerLiteralNames = []string{
	"", "','", "'{'", "'}'", "':'", "'['", "']'", "'='", "'||'", "'&&'", "'.'",
	"';'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'+'", "'-'", "'*'",
	"'/'", "'%'", "'!'", "'?'", "'??'", "'as?'", "'('", "')'", "'transaction'",
	"'struct'", "'interface'", "'fun'", "'pre'", "'post'", "'pub'", "'pub(set)'",
	"'return'", "'break'", "'continue'", "'let'", "'var'", "'if'", "'else'",
	"'while'", "'true'", "'false'", "'nil'", "'import'", "'from'",
}

var lexerSymbolicNames = []string{
	"", "", "", "", "", "", "", "", "", "", "", "", "Equal", "Unequal", "Less",
	"Greater", "LessEqual", "GreaterEqual", "Plus", "Minus", "Mul", "Div",
	"Mod", "Negate", "Optional", "NilCoalescing", "FailableDowncasting", "OpenParen",
	"CloseParen", "Transaction", "Struct", "Interface", "Fun", "Pre", "Post",
	"Pub", "PubSet", "Return", "Break", "Continue", "Let", "Var", "If", "Else",
	"While", "True", "False", "Nil", "Import", "From", "Identifier", "DecimalLiteral",
	"BinaryLiteral", "OctalLiteral", "HexadecimalLiteral", "InvalidNumberLiteral",
	"StringLiteral", "WS", "Terminator", "BlockComment", "LineComment",
}

var lexerRuleNames = []string{
	"T__0", "T__1", "T__2", "T__3", "T__4", "T__5", "T__6", "T__7", "T__8",
	"T__9", "T__10", "Equal", "Unequal", "Less", "Greater", "LessEqual", "GreaterEqual",
	"Plus", "Minus", "Mul", "Div", "Mod", "Negate", "Optional", "NilCoalescing",
	"FailableDowncasting", "OpenParen", "CloseParen", "Transaction", "Struct",
	"Interface", "Fun", "Pre", "Post", "Pub", "PubSet", "Return", "Break",
	"Continue", "Let", "Var", "If", "Else", "While", "True", "False", "Nil",
	"Import", "From", "Identifier", "IdentifierHead", "IdentifierCharacter",
	"DecimalLiteral", "BinaryLiteral", "OctalLiteral", "HexadecimalLiteral",
	"InvalidNumberLiteral", "StringLiteral", "QuotedText", "EscapedCharacter",
	"HexadecimalDigit", "WS", "Terminator", "BlockComment", "LineComment",
}

type StrictusLexer struct {
	*antlr.BaseLexer
	channelNames []string
	modeNames    []string
	// TODO: EOF string
}

var lexerDecisionToDFA = make([]*antlr.DFA, len(lexerAtn.DecisionToState))

func init() {
	for index, ds := range lexerAtn.DecisionToState {
		lexerDecisionToDFA[index] = antlr.NewDFA(ds, index)
	}
}

func NewStrictusLexer(input antlr.CharStream) *StrictusLexer {

	l := new(StrictusLexer)

	l.BaseLexer = antlr.NewBaseLexer(input)
	l.Interpreter = antlr.NewLexerATNSimulator(l, lexerAtn, lexerDecisionToDFA, antlr.NewPredictionContextCache())

	l.channelNames = lexerChannelNames
	l.modeNames = lexerModeNames
	l.RuleNames = lexerRuleNames
	l.LiteralNames = lexerLiteralNames
	l.SymbolicNames = lexerSymbolicNames
	l.GrammarFileName = "Strictus.g4"
	// TODO: l.EOF = antlr.TokenEOF

	return l
}

// StrictusLexer tokens.
const (
	StrictusLexerT__0                 = 1
	StrictusLexerT__1                 = 2
	StrictusLexerT__2                 = 3
	StrictusLexerT__3                 = 4
	StrictusLexerT__4                 = 5
	StrictusLexerT__5                 = 6
	StrictusLexerT__6                 = 7
	StrictusLexerT__7                 = 8
	StrictusLexerT__8                 = 9
	StrictusLexerT__9                 = 10
	StrictusLexerT__10                = 11
	StrictusLexerEqual                = 12
	StrictusLexerUnequal              = 13
	StrictusLexerLess                 = 14
	StrictusLexerGreater              = 15
	StrictusLexerLessEqual            = 16
	StrictusLexerGreaterEqual         = 17
	StrictusLexerPlus                 = 18
	StrictusLexerMinus                = 19
	StrictusLexerMul                  = 20
	StrictusLexerDiv                  = 21
	StrictusLexerMod                  = 22
	StrictusLexerNegate               = 23
	StrictusLexerOptional             = 24
	StrictusLexerNilCoalescing        = 25
	StrictusLexerFailableDowncasting  = 26
	StrictusLexerOpenParen            = 27
	StrictusLexerCloseParen           = 28
	StrictusLexerTransaction          = 29
	StrictusLexerStruct               = 30
	StrictusLexerInterface            = 31
	StrictusLexerFun                  = 32
	StrictusLexerPre                  = 33
	StrictusLexerPost                 = 34
	StrictusLexerPub                  = 35
	StrictusLexerPubSet               = 36
	StrictusLexerReturn               = 37
	StrictusLexerBreak                = 38
	StrictusLexerContinue             = 39
	StrictusLexerLet                  = 40
	StrictusLexerVar                  = 41
	StrictusLexerIf                   = 42
	StrictusLexerElse                 = 43
	StrictusLexerWhile                = 44
	StrictusLexerTrue                 = 45
	StrictusLexerFalse                = 46
	StrictusLexerNil                  = 47
	StrictusLexerImport               = 48
	StrictusLexerFrom                 = 49
	StrictusLexerIdentifier           = 50
	StrictusLexerDecimalLiteral       = 51
	StrictusLexerBinaryLiteral        = 52
	StrictusLexerOctalLiteral         = 53
	StrictusLexerHexadecimalLiteral   = 54
	StrictusLexerInvalidNumberLiteral = 55
	StrictusLexerStringLiteral        = 56
	StrictusLexerWS                   = 57
	StrictusLexerTerminator           = 58
	StrictusLexerBlockComment         = 59
	StrictusLexerLineComment          = 60
)
