package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := `1VAR-1
1 h4    Kc4
2 h5    Kd5
3 h6    Ke6
4 h7    Kf7
5 h8=Q
White win teaches
how to count!

2VAR-1
1 a4     Ke4 
2 a5   Kd5 
3 a6    Kc6 
4 a7    Kb7
White loses Pawn
or Queen for draw.

2VAR-2
1 Kf5!    Ke3
2 a4    Kd4
3 a5    Kc5
4 a6    Kb6
White King’s lost
tempo and Pawn
draws.

2VAR-3
1 Kf5!    Ke3
2 Ke5!    Kd3
3 a4    Kc4
4 a5    Kb5
and Black draws
with cut-in.

2VAR-4
1 Kf5!   Ke3
2 Ke5!  Kd3
3 Kd5!  Kc3
4 a4       Kb4
nails down Pawn
and draws.   

2VAR-5
1 Kf5!  Ke3
2 Ke5!  Kd3
3 Kd5!  Kc3
4 Kc5!  Kd3
5 a4     Kc3
6 a5     Kb3
7 a6     Ka4
8 a7     Ka5
too late Black, loses  
with  9 a8=Q#.

3VAR-1
1 e6    Kd8 
2 e7+    Ke8 
3 Ke6
to protect the Pawn
but stalemates Black.

3VAR-2
1 Ke6! Kf8 
2 Kd7     Kf7 
3 e6+     Kf8  
4 e7+
White wins.

4VAR-1
1 Kf5    Kf8
2 Ke6    Ke8
3 d7+    Kd8
4 Kd6
and Black draws by stalemate.

5VAR-1
1 Kd5    Kd7
2 Kc5    Kd8
3 Kc6    Kc8 
4 b7+    Kb8
5 Kb6
draws.

6VAR-1
1 c5    Kf7
2 Kg5    Ke6
Black pinches off King,
wins the Pawn and draws.

6VAR-2
1 Kg6!     Kf8
2 Kf6     Ke8
3 Ke6     Kd8
4 Kd6     Kc8
5 Kc6      Kd8
6 Kb7     Ke7
7 c5
White wins when Pawn 
cannot be stopped. 

7VAR-1
1 g6+    Kh8   
2 Kf7
stalemates Black.

7VAR-2
1 g6+    Kh8   
2 g7+    Kg8 
3 Kg6
 Black is stalemated.

8VAR-1
1 Ke4    Ke6
2 e3!    Kd6
3 Kf5    Kd5
4 e4+    Kd6
5 Kf6!    Kd7
6 e5    Ke8
7 Ke6
familiar opposition
and White wins.

8VAR-2
1 Ke4    Ke6
2 e3!    Kd6
3 Kf5    Ke7
4 Ke5    Kd7
5 Kf6    Kd6
6 e4    Kd7
7 e5    Ke8
8 Ke6
White wins.

8VAR-3
1 Ke4    Ke6
2 e3!    Kd6
3 Kf5    Ke7
4 Ke5    Kd7
5 Kf6    Kd6
6 e4    Kd7
7 e5    Kd8
8 Kf7
White wins.

9VAR-1
1 Kf2    Kg6
2  Ke3    Kf5
3  Kd4!    Ke6
4  Kc5    Ke5
5 d4+    Ke6
6 d5+    Kd7
7 d6    Kd8
8 Kc6    Kc8
9 d7+    Kd8
 and Black draws 

9VAR-2
1 Kf2    Kg6
2  Ke3    Kf5
3  Kd4!    Ke6
4  Kc5    Ke5
5 d4+    Ke6
6 Kc6    Ke7
7 d5    Ke8
7 Kd6
with familiar White 
opposition and win.

9VAR-3
1 Kf2    Kg6
2  Ke3    Kf5
3  Kd4!    Ke6
4  Kc5    Kd7
5 Kd5 Ke7
6 d4    Kd7
Black has opposition 
this time and draws.

9VAR-4
1 Kf2     Kg6
2  Ke3     Kf5
3  Kd4!     Ke6
4  Kc5     Kd7
5 Kd5  Ke7
6  Kc6     Kd8
7 Kd6
and White wins.

9VAR-5
1 Kf2    Kg6
2 Ke3    Kf5
3 Kd4!    Ke6
4 Kc5    Kd7
5 Kd5 Ke7
6 Kc6     Ke6
7 d4     Ke7
8 d5    Ke8
9 Kc7 
ends it (for Black). 

10VAR-1
1 g4    Kc5 
2 g5    Kd5
3 g6    Ke6
4 g7    Kf7
5 g8=Q Kxg8
Pawn is lost 
leaving a draw.

10VAR-2
1 Kb3    Kc5
2 Kc3    Kd5
3 Kd3    Ke5
4 Ke3    Kf5
5 Kf3    Kg5
6 Kg3    Kf5
7 Kh4    Kf6
8 Kh5    Kf5
9 g4+    Kf6
10 Kh6      Kf7
11 g5    Kf8
12 Kh7
White wins

10VAR-3
1 Kb3    Kc5
2 Kc3    Kd5
3 Kd3    Ke5
4 Ke3    Kf5
5 Kf3    Kg5
6 Kg3    Kf5
7 Kh4    Kf6
8 Kh5    Kf5
9 g4+    Kf6
10 Kh6      Kf7
11 g5    Kf8
12 Kh7 
White wins again.

10VAR-4
1 Kb3    Kc5
2 Kc3    Kd5
3 Kd3    Ke5
4 Ke3    Kf5
5 Kf3    Kg5
6 Kg3    Kf5
7 Kh4    Kf6
8 Kh5    Kf5
9 g4+    Kf6
10 Kh6      Kf7
11 g5    Kg8
12 Kg6 
White wins.

10VAR-5
1. Kb3 Kc5
2. Kc3 Kd5
3. Kd3 Ke5
4. Ke3 Kf5
5. Kf3 Kg5
6. Kg3 Kf5
7. Kh4 Kf6
8. Kh5 Kg7
9. Kg5 Kf7
10. Kh6 Kf6
11. g4 Kf7
12. g5
White wins again.

11VAR-1
1 Kb6    Kc8
2 a4    Kb8
3 a5    Ka8
4 a6    Kb8
5 a7+    Ka8
6 Ka6
 Black is stalemated.

11VAR-2
1 Kb6    Kc8
2 Ka7    Kc7
3 a4    Kc8
4 a5    Kc7
5 a6    Kc8
6 Ka8    Kc7
7 a7    Kc8
 White is stalemated.

12VAR-1
1  Kf 4    Kg7
2  Kf5    Kh8
3 Ke6    Kg7
4 Kd7    Kh8
King must not come
any closer -- stalemate.

13VAR-1
1 g7    Kh7
2 Kf7
allows a draw by stalemate.

13VAR-2
1 g7    Kh7
2 g6+    Kg8
 lets the win slip 
forever-- stalemate.

13VAR-3
1 g7    Kh7
2 g8=Q+!  Kxg8
3 Kg6     Kf8
4 Kh7
White wins.

14VAR-1
1 a4!     Kxc3
2 a5
Pawn cannot be 
overtaken.

14VAR-2
1 a4!    Kc5
2 c4    Kxc4
Pawn is captured 
with a draw.    

14VAR-3
1 a4!    Kc5
2 a5    Kb5
Pawns are captured 
with a draw.

14VAR-4
1 a4!     Kc5
2 Kg3    Kc4
3 Kf3
one Pawn is captured. 
but parries the other.   

14VAR-5
1 a4!     Kc5
2 Kg3    Kb6
3 c4    Kc5
4 a5    
and White’s 
Pawn Queens.

14VAR-6
1 a4!     Kc5
2 Kg3    Kb6
3 c4    Ka5
4 c5    Ka6
5 Kf3    Kb7
6 a5    Kc6
7 a6    Kc7
8 Ke4    Kc6
9 Ke5    Kc7
10 Kd5    Kb8
11 c6    Ka7
12 c7
White wins. 

15VAR-1
1 Kf4    g5+
2 Kf5    g4
3 Ke6    g3
4 f7     g2
5 f8=Q+
White takes Pawn
or Queen and wins.

15VAR-2
1 Kf4    g5+
2 Kf5    g4
3 Ke6    g3
4 f7    Kg7
5 Ke7
White wins.

15VAR-3
1 Kf4    Kh7
2 Kg5    Kh8
3 Kh6!    g5
4 f7
followed by 
f8=Q#.

16VAR-1
1 Kc3    a3! 
2 bxa3    Ke6
Black draws easily.

16VAR-2
1 Kc3    a3!
2 b4    Ke5
3 Kb3    Kd5
4 Kxa3    Kc6
5 Ka4    Kb6
6 b5    Kb7
7 Ka5    Ka7
8 b6+    Kb7
9 Kb5    Kb8
10 Kc6    Kc8
White cannot 
avoid the draw.

16VAR-3
1 Kb1!    Ke5
2 Ka2    Kd5
3 Ka3    Kc5
4 Kxa4     Kb6
5 Kb4
 White having the
 opposition wins. 

16VAR-4
1 Kb1!     a3
2   b3!    Ke5
3   Ka2    Kd5
4   Kxa3    Kc6
5 Kb4!    Kb6
Black with the
opposition draws.

17VAR-1
1 a4    Ke4
2 a5    Kd5
3 a6    Kd6
4 a7    Kc7
5 a8=Q
White wins

18VAR-1
1 Ke7     Kc3
2 Kd7     Kd4
3 Kc7     Kc5
4 Kb7     Kd6
5 Kxa7     Kc7
6 Ka8     Kc8
and Black gets
a draw.

19VAR-1
1 b5    h4
2 b6    h3
3 b7    h2
4 b8=Q+
with Pawn loss
Black loses game.

19VAR-2
1 b5    Ke5
2 Kc5     Ke6
3 b6     Kd7
is a draw.

19VAR-3
1 b5    Ke5
2 Kc5     Ke6
3 Kc6     h4
4 b6    h3
5 b7    h2
6 b8=Q h1=Q
Position is a draw.

19VAR-4
1 b5    Ke5
2 b6!    Kd6
3 Kb5    Kd7
4 Ka6    Kc8
5 Ka7
White wins.

19VAR-5
1 b5    Ke5
2 b6!    Kd6
3 Kb5    h4
4 b7    Kc7
5 Ka6    Kb8
and White loses. 

20VAR-1
1 e6    g3
2 e7    g2
3 e8=Q g1=Q
4 Qg8+
Black’s Queen
comes off the board.

20VAR-2
1 e6    Kf6
2 Kd6    g3
3 e7    Kf7
4 Kd7
threat of Queening with
check wins for White.

21VAR-1
1 Kg5    b5
2 h5    b4
 both Pawns Queen
 leading to a draw.

21VAR-2
1 Kg5    b5
2 Kf4    b4
3 Ke4 
catching the 
Black Pawn.

21VAR-3
1 Kg5    b5
2 Kf4    Ke2
3 Ke4    b4
Pawn is caught,
White wins.

21VAR-4
1 Kg5    b5
2 Kf4    Ke2
3 Ke4    Kd2
4 Kd4    Kc2
5 h5    b4
6 h6    b3
7 h7    b2
8 h8=Q b1=Q 
Black draws. 

21VAR-5
1 Kg5    b5
2 Kf4    Ke2
3 Ke4    Kd2
4 Kd4    Kc2
5 Kc5    Kc3
6 Kxb5    Kd4
7 h5    Ke5
White Pawn’s loss 
Is a smashing draw.

22VAR-1
1 f7    h2
2 f8=Q    h1=Q
3 Qf3+    Kh2
allows 4 Qg3#.

23VAR-1
1 g4    b5
2 g5    b4
3 g6    b3+
4 Kc3    b2
5 g7    b1=Q
6 g8=Q+ Ka1
   and a draw.

23VAR-2
1 Kc3    b5
2 Kb4    Kb2
3 g4
White wins easily.

23VAR-3
1 Kc3    Ka3
2 Kc4    Ka4
3 g4    b5+
4 Kd3!    Ka3
5 g5    b4
6 g6    b3
7 g7    b2
8 Kc2!    Ka2
9 g8=Q+ Ka3
White mates 
with 10 Qb3#. 

24VAR-1
1 Kc5!    Kg6
2 b4    Kf7
3 b5    Ke7
4 Kc6    Kd8
5 Kb7    g5
6 b6    g4
7 Ka7    g3
8 b7    g2
9 b8=Q+
   wins for White.

24VAR-2
1 Kc5!    g5
2 b4    g4
3 Kd4    g3
4 Ke3    Kg5
5 Kf3     Kh4
6 Kg2
White’s Pawn will
Queen and win. 

24VAR-3
1 Kc5!    g5
2 b4    g4
3 Kd4    g3
4 Ke3    Kg5
5 Kf3     Kf5
6 Kxg3     Ke4
Black captures 
Pawn and draws.

25VAR-1
1 Kxb7     Kb3
2 f4    Kc4
3 f5    Kd5
both lose Pawns
and draw.

25VAR-2
1 Kxb7    Kb3 
2 Kc6    Kc4 
3 Kd6    Kd4
4 f4    Ke4
Black also wins 
Pawn for draw. 

25VAR-3
1 f4    b5 
2 f5    b4
both sides Queen 
and draw. 

25VAR-4
1 Kd6!     b5
2 Kc5     Kb3 
3 Kxb5     Kc3 
4 Kc5     Kd3 
5 Kd5 
and the Pawn 
is safe.

25VAR-5
1 Kd6!     Ka3
2 Kc5     Ka4
3 f4     b5
4 f5     b4
5 f6     b3 
6 f7     b2
7 f8=Q     b1=Q 
8 Qa8+     Kb3 
9 Qb7+ Kc2
is only a draw.

25VAR-6
1 Kd6!    Ka3
2 Kc5    Ka4
3 f4    b5
4 f5    b4
5 Kc4!    b3
6 Kc3!    Ka3
7 f6    b2
8 f7    b1=Q
9 f8=Q+  Ka2
and mate next 
with 10 Qa8#.
`

	initialBoards := map[string]string{
		"1":   "8/8/8/8/8/1k5P/8/2K5 w - - 0 1",
		"2":   "8/8/8/6K1/8/5k2/P7/8 w - - 0 1",
		"3":   "4k3/8/3K4/4P3/8/8/8/8 w - - 0 1",
		"4":   "8/5k2/3P4/8/6K1/8/8/8 w - - 0 1",
		"5":   "3k4/8/1P6/8/4K3/8/8/8 w - - 0 1",
		"6":   "6k1/8/7K/8/2P5/8/8/8 w - - 0 1",
		"7":   "8/7k/5K2/6P1/8/8/8/8 w - - 0 1",
		"8":   "8/8/5k2/8/5K2/8/4P3/8 w - - 0 1",
		"9":   "8/7k/8/8/8/3P4/8/6K1 w - - 0 1",
		"10":  "8/8/8/1k6/8/K7/6P1/8 w - - 0 1",
		"11":  "3k4/8/K7/8/8/8/P7/8 w - - 0 1",
		"12":  "7k/7P/6P1/8/8/6K1/8/8 w - - 0 1",
		"13":  "6k1/8/5KP1/6P1/8/8/8/8 w - - 0 1",
		"14":  "8/8/8/8/2k5/P1P5/7K/8 w - - 0 1",
		"15":  "8/8/5Ppk/8/8/4K3/8/8 w - - 0 1",
		"16":  "8/8/5k2/8/p7/8/1PK5/8 w - - 0 1",
		"17":  "8/8/2p5/8/8/P4k2/8/2K5 w - - 0 1",
		"18":  "8/p4K2/P7/8/8/8/1k6/8 w - - 0 1",
		"19":  "8/8/8/7p/1PK2k2/8/8/8 w - - 0 1",
		"20":  "8/8/8/3KP1k1/6p1/8/8/8 w - - 0 1",
		"21":  "8/1p6/7K/8/7P/8/5k2/8 w - - 0 1",
		"22":  "8/8/5P2/8/6K1/7p/6k1/8 w - - 0 1",
		"23":  "8/8/1p6/8/8/6P1/k1K5/8 w - - 0 1",
		"24":  "8/6p1/7k/8/1K6/8/1P6/8 w - - 0 1",
		"25":  "8/1pK5/8/8/8/8/k4P2/8 w - - 0 1",
		"26":  "8/2p5/6K1/8/8/5k2/P7/8 w - - 0 1",
		"27":  "8/4K1pp/8/8/8/8/k6P/8 w - - 0 1",
		"28":  "8/6k1/1p6/6KP/P7/8/8/8 w - - 0 1",
		"29":  "8/8/5k2/p7/P3KP2/8/8/8 w - - 0 1",
		"30":  "8/2p1k3/2P5/1P2K3/8/8/8/8 w - - 0 1",
		"31":  "8/kp6/8/PP6/1K6/8/8/8 w - - 0 1",
		"32":  "8/6k1/3p4/3P1PK1/8/8/8/8 w - - 0 1",
		"33":  "8/3k4/2pP4/2P1K3/8/8/8/8 w - - 0 1",
		"34":  "8/1k6/2pK4/8/1P1P4/8/8/8 w - - 0 1",
		"35":  "2k5/1p6/1P6/2PK4/8/8/8/8 w - - 0 1",
		"36":  "8/8/kpK5/8/P1P5/8/8/8 w - - 0 1",
		"37":  "3k4/2p5/2K5/1P1P4/8/8/8/8 w - - 0 1",
		"38":  "4k3/8/4KP2/7p/8/6P1/8/8 w - - 0 1",
		"39":  "6k1/7p/7K/7P/8/8/6P1/8 w - - 0 1",
		"40":  "8/8/8/5kP1/4p2P/4K3/8/8 w - - 0 1",
		"41":  "8/2k5/p1P5/P1K5/8/8/8/8 w - - 0 1",
		"42":  "8/1pk5/8/1K6/1PP5/8/8/8 w - - 0 1",
		"43":  "8/8/8/8/4k1p1/8/4KPP1/8 w - - 0 1",
		"44":  "8/8/1kp5/8/K1PP4/8/8/8 w - - 0 1",
		"45":  "8/p7/P7/8/8/4k3/4P3/4K3 w - - 0 1",
		"46":  "3k4/1K3p2/8/4P3/8/6P1/8/8 w - - 0 1",
		"47":  "8/6p1/8/5P2/5P2/5K2/8/6k1 w - - 0 1",
		"48":  "8/6p1/8/K6P/7P/1k6/8/8 w - - 0 1",
		"49":  "8/6p1/8/4K3/7P/4k2P/8/8 w - - 0 1",
		"50":  "5k2/5p2/8/4P3/2K1P3/8/8/8 w - - 0 1",
		"51":  "4K3/8/2p5/8/P2k4/8/P7/8 w - - 0 1",
		"52":  "8/8/8/8/1k6/4p1P1/6P1/6K1 w - - 0 1",
		"53":  "8/6p1/8/8/7K/7P/4k2P/8 w - - 0 1",
		"54":  "8/8/3k4/2p2p2/P1K2P2/8/8/8 w - - 0 1",
		"55":  "8/p5kp/8/1P6/8/1K6/P7/8 w - - 0 1",
		"56":  "8/2k5/2Pp3p/1P6/8/5K2/8/8 w - - 0 1",
		"57":  "8/7k/4K2p/6p1/6P1/7P/8/8 w - - 0 1",
		"58":  "7k/6p1/6P1/8/8/p5K1/P7/8 w - - 0 1",
		"59":  "8/8/8/7p/5Ppk/4KP2/8/8 w - - 0 1",
		"60":  "8/8/8/p2K1p2/3P1Pk1/8/8/8 w - - 0 1",
		"61":  "8/5k2/5p2/p7/2P1K1P1/8/8/8 w - - 0 1",
		"62":  "2K5/8/5p2/3k2p1/P5P1/8/8/8 w - - 0 1",
		"63":  "8/8/7p/p7/k1P2K2/8/P7/8 w - - 0 1",
		"64":  "8/8/8/KP6/1p6/k4p2/5P2/8 w - - 0 1",
		"65":  "8/8/4k3/1p5p/1P5P/4K3/8/8 w - - 0 1",
		"66":  "4K3/8/8/1p5p/1P5P/8/8/4k3 w - - 0 1",
		"67":  "8/4p2p/8/1K1k4/8/5P2/P7/8 w - - 0 1",
		"68":  "8/8/5p2/6p1/P2k2P1/8/8/K7 w - - 0 1",
		"69":  "8/p2p4/8/8/8/k7/5P1P/7K w - - 0 1",
		"70":  "8/8/2k2ppp/8/P2K2PP/8/8/8 w - - 0 1",
		"71":  "8/p6p/4k3/8/4K3/8/P5PP/8 w - - 0 1",
		"72":  "8/8/8/1p3k1p/1P5P/5KP1/8/8 w - - 0 1",
		"73":  "8/3pp2p/8/1kPP4/5P2/8/1K6/8 w - - 0 1",
		"74":  "8/4k1pp/8/4KPPP/8/8/8/8 w - - 0 1",
		"75":  "8/p7/Pp6/1P6/8/K2k4/P7/8 w - - 0 1",
		"76":  "k7/P7/1P6/4p3/4Pp2/5K2/8/8 w - - 0 1",
		"77":  "8/8/8/2p5/2Pp4/3K2Pk/7P/8 w - - 0 1",
		"78":  "K7/5p2/k7/6p1/P5P1/8/P7/8 w - - 0 1",
		"79":  "8/1p6/p3k3/8/6p1/PP6/4KPP1/8 w - - 0 1",
		"80":  "8/1pp5/p2p4/P2P4/1PP5/7k/8/7K w - - 0 1",
		"81":  "3k4/2p5/1pKp4/p2P4/2P5/P7/1P6/8 w - - ",
		"82":  "7k/8/5PpK/Pp1P2pp/3P4/8/5p2/8 w - - 0 1",
		"83":  "k7/2p1pp2/2P3p1/4P1P1/5P2/p7/Kp3P2/8 w ",
		"84":  "8/2p5/1pPp4/1P1Pp3/4Pp1k/5P2/5KP1/8 w ",
		"85":  "8/2pp2pp/8/2PP1P2/1p5k/8/PP4p1/6K1 w - ",
		"86":  "3K4/8/8/1n5P/5k2/8/8/8 w - - 0 1",
		"87":  "8/8/8/4P3/2k5/6K1/8/2n5 w - - 0 1",
		"88":  "8/8/1P6/8/2K2kn1/8/6P1/8 w - - 0 1",
		"89":  "8/5n2/8/3P2PK/3k2p1/8/8/8 w - - 0 1",
		"90":  "8/7K/7P/1P6/2n4k/4n3/1P6/8 w - - 0 1",
		"91":  "8/6b1/5k2/8/P3K1P1/8/8/8 w - - 0 1",
		"92":  "8/8/1P6/8/2K2k2/8/1b4P1/8 w - - 0 1",
		"93":  "8/2b5/7P/1p6/1P4K1/8/8/5k2 w - - 0 1",
		"94":  "6b1/8/2k4K/4P3/7P/8/8/8 w - - 0 1",
		"95":  "8/8/5b2/5P2/2P1K1k1/8/8/8 w - - 0 1",
		"96":  "4k3/8/2K5/8/PP6/8/8/4b3 w - - 0 1",
		"97":  "K7/8/kPP5/1b6/1P6/8/8/8 w - - 0 1",
		"98":  "8/5b2/7p/7P/5PP1/8/1K4k1/8 w - - 0 1",
		"99":  "4b1k1/4P1P1/5P1K/8/8/8/8/8 w - - 0 1",
		"100": "6k1/3p4/P4P2/2P5/1K6/7b/8/8 w - - 0 1",
		"101": "8/8/1KP5/3r4/8/8/8/k7 w - - 0 1",
		"102": "6K1/1k4P1/6P1/8/8/8/r7/8 w - - 0 1",
		"103": "8/8/8/8/pN6/8/2K5/k7 w - - 0 1",
		"104": "8/3N4/8/8/8/2P1k3/8/7K w - - 0 1",
		"105": "8/8/8/8/2K5/p1N5/Pk6/8 w - - 0 1",
		"106": "8/5k2/7P/6K1/1N6/p7/8/8 w - - 0 1",
		"107": "8/8/4N3/8/7p/3K1k2/7P/8 w - - 0 1",
		"108": "8/8/7p/8/6pp/4N1Pk/5K2/8 w - - 0 1",
		"109": "b3K3/8/3P4/4k1N1/8/8/8/8 w - - 0 1",
		"110": "8/Kb6/1Pk5/8/7N/8/8/8 w - - 0 1",
		"111": "8/8/8/1PK5/3N4/8/5kb1/8 w - - 0 1",
		"112": "8/P1K1k3/8/8/2N5/5b2/8/8 w - - 0 1",
		"113": "8/3PK3/8/N1k5/5b2/8/8/8 w - - 0 1",
		"114": "4k3/8/4N3/5n1P/8/8/8/K7 w - - 0 1",
		"115": "8/7k/8/2K5/1P2N3/n7/8/8 w - - 0 1",
		"116": "K2n4/8/P7/8/2k5/8/2N5/8 w - - 0 1",
		"117": "8/KP1n4/2k5/8/5N2/8/8/8 w - - 0 1",
		"118": "1K6/P7/1nk5/8/4N3/8/8/8 w - - 0 1",
		"119": "k4n2/8/8/2N1P3/5K2/8/8/8 w - - 0 1",
		"120": "8/8/8/5P1p/3N3k/8/8/3n3K w - - 0 1",
		"121": "8/8/2K5/kp2p3/p4P2/3N4/8/3n4 w - - 0 1",
		"122": "8/pK6/pP6/8/8/r6N/8/5k2 w - - 0 1",
		"123": "7k/4K1p1/6Pp/5N1P/8/8/8/8 w - - 0 1",
		"124": "1K6/5p2/k3P3/p7/8/3p4/P1N5/8 w - - 0 1",
		"125": "2n5/8/1PN1Pk2/8/4K3/8/8/8 w - - 0 1",
		"126": "5n2/8/3K4/2N2k1P/8/6P1/8/8 w - - 0 1",
		"127": "5n2/8/7P/4k1P1/p4N2/8/5K2/8 w - - 0 1",
		"128": "8/8/2k4b/P7/8/8/2N2PKp/8 w - - 0 1",
		"129": "4k3/3p1N2/8/P1P4b/8/8/7K/8 w - - 0 1",
		"130": "8/1pN2p2/1P6/8/2P5/5k2/b7/6K1 w - - 0 1",
		"131": "8/8/P4PK1/8/6N1/8/8/r6k w - - 0 1",
		"132": "8/7p/1P5K/3N1r2/8/3k4/6P1/8 w - - 0 1",
		"133": "1k2N2b/2p4P/2p2p2/2P2P2/8/5K2/8/8 w - ",
		"134": "2K5/p7/k1B5/p7/p7/8/8/8 w - - 0 1",
		"135": "8/4k3/8/7P/4B3/5K2/8/8 w - - 0 1",
		"136": "4k3/8/8/7P/8/4K2B/8/8 w - - 0 1",
		"137": "6K1/8/7p/8/k7/2B5/1P6/8 w - - 0 1",
		"138": "8/3k4/P7/K7/1B6/1p6/8/8 w - - 0 1",
		"139": "6k1/3p4/8/8/8/B7/P7/7K w - - 0 1",
		"140": "8/8/k1K5/8/pB6/P7/8/8 w - - 0 1",
		"141": "8/2P5/1K6/3k4/8/p7/6p1/B7 w - - 0 1",
		"142": "8/3p4/8/7p/1KPk4/8/B7/8 w - - 0 1",
		"143": "8/8/8/1pk1K3/2p5/8/1P5B/8 w - - 0 1",
		"144": "8/k7/8/1pK5/p4B2/P7/8/8 w - - 0 1",
		"145": "8/8/3p4/4p2B/4P3/8/4K1kp/8 w - - 0 1",
		"146": "8/3K4/3P1p2/p4p2/k3B3/p7/8/8 w - - 0 1",
		"147": "3n4/2K5/1P6/k7/2B5/8/8/8 w - - 0 1",
		"148": "8/8/3B4/8/7K/5k1P/8/3n4 w - - 0 1",
		"149": "6K1/5P2/6k1/2b5/8/8/3B4/8 w - - 0 1",
		"150": "3B4/5K2/4P3/2bk4/8/8/8/8 w - - 0 1",
		"151": "8/4B1b1/6P1/5K1k/8/8/8/8 w - - 0 1",
		"152": "5k2/8/6PK/8/8/2b3B1/8/8 w - - 0 1",
		"153": "8/8/6K1/4B2P/6k1/4b3/8/8 w - - 0 1",
		"154": "2KB4/1P6/2k5/8/8/8/7b/8 w - - 0 1",
		"155": "2k5/P1p5/1n2K3/8/4B3/8/8/8 w - - 0 1",
		"156": "2B5/8/5k1K/4p3/1b4P1/8/8/8 w - - 0 1",
		"157": "1B6/8/7P/4p3/3b3k/8/8/2K5 w - - 0 1",
		"158": "6b1/5p2/8/2B2K2/7k/8/6P1/8 w - - 0 1",
		"159": "b7/6K1/2p5/8/4Pp2/8/2k3B1/8 w - - 0 1",
		"160": "6k1/8/p4K2/2P5/6p1/1b6/8/5B2 w - - 0 1",
		"161": "k7/8/3P2K1/B7/8/8/7r/8 w - - 0 1",
		"162": "8/8/1P6/3K4/1k6/5rB1/8/8 w - - 0 1",
		"163": "8/8/2P5/4K3/8/1B5k/6r1/8 w - - 0 1",
		"164": "8/8/P7/3p3r/3k3B/3P4/5K2/8 w - - 0 1",
		"165": "4k3/8/1P2b3/4K3/4B3/8/8/r7 w - - 0 1",
		"166": "8/3b1K1B/6Pk/8/8/8/7P/8 w - - 0 1",
		"167": "5k2/8/5K2/6P1/5P2/1bB5/8/8 w - - 0 1",
		"168": "5b2/5k2/8/4P3/5P2/5K2/4B3/8 w - - 0 1",
		"169": "8/8/8/4bk2/8/8/3B1PP1/5K2 w - - 0 1",
		"170": "8/8/7k/8/8/6Pb/4B2P/6K1 w - - 0 1",
		"171": "8/5b2/8/4k3/4B3/3PKP2/8/8 w - - 0 1",
		"172": "4k3/4P3/8/4Pp2/7K/8/4B2n/8 w - - 0 1",
		"173": "4b2k/7p/1p6/6K1/2P1B1P1/8/8/8 w - - 0 1",
		"174": "k2K4/8/6P1/2B5/5P2/p7/4p2b/8 w - - 0 1",
		"175": "K7/8/Pp6/kp6/1n6/1P6/8/B7 w - - 0 1",
		"176": "6n1/7k/2p2p2/8/P2P4/8/8/2B4K w - - 0 1",
		"177": "8/8/2r3pP/6P1/8/K7/7k/3B4 w - - 0 1",
		"178": "1kr5/4P1Pp/8/8/8/B7/7K/8 w - - 0 1",
		"179": "4k3/4p3/P2p4/8/2bP4/p7/2P5/K2B4 w - - ",
		"180": "8/2p2b2/Pp1p4/4pp2/k7/2P5/1P2BPK1/8 w ",
		"181": "8/1K6/8/8/3pk3/8/8/R7 w - - 0 1",
		"182": "K6R/8/7p/6k1/8/8/8/8 w - - 0 1",
		"183": "K5R1/8/6p1/5k2/8/8/8/8 w - - 0 1",
		"184": "8/7p/8/8/6k1/8/8/KR6 w - - 0 1",
		"185": "1R6/8/8/7p/6k1/8/8/K7 w - - 0 1",
		"186": "8/5K2/R7/8/p7/k7/8/8 w - - 0 1",
		"187": "3R4/4K3/8/5kp1/8/8/8/8 w - - 0 1",
		"188": "2R5/7k/5K2/8/p7/8/1p6/8 w - - 0 1",
		"189": "K7/2k5/R7/8/5p2/6p1/8/8 w - - 0 1",
		"190": "8/8/8/3R4/2K5/6pp/k7/8 w - - 0 1",
		"191": "8/8/p4K1R/8/p7/8/8/k7 w - - 0 1",
		"192": "6k1/8/4K3/3R4/3p4/8/7b/8 w - - 0 1",
		"193": "8/p7/8/1P6/8/3K4/p3R3/1k6 w - - 0 1",
		"194": "8/kp6/4R3/1P2K3/8/3pp3/8/8 w - - 0 1",
		"195": "8/8/8/pp5R/k6p/2K5/1P4p1/8 w - - 0 1",
		"196": "5Rn1/7b/8/8/k5P1/8/3K4/8 w - - 0 1",
		"197": "8/8/4P2R/8/1r6/k7/8/K7 w - - 0 1",
		"198": "8/8/8/4kr1P/7K/8/8/7R w - - 0 1",
		"199": "3r4/8/8/8/k2P4/3K4/8/1R6 w - - 0 1",
		"200": "K7/P5R1/2k5/8/8/8/2r5/8 w - - 0 1",
		"201": "R7/P7/8/8/6K1/8/6k1/r7 w - - 0 1",
		"202": "K7/P3k3/8/8/8/8/2R5/1r6 w - - 0 1",
		"203": "R7/1K1k4/P7/8/8/8/8/2r5 w - - 0 1",
		"204": "r7/8/4k3/P7/K7/3R4/8/8 w - - 0 1",
		"205": "8/r4k2/8/8/P7/K3R3/8/8 w - - 0 1",
		"206": "6K1/6R1/r5P1/8/7k/8/8/8 w - - 0 1",
		"207": "7r/6k1/1P6/8/8/5R2/8/K7 w - - 0 1",
		"208": "1r6/4k3/8/8/1P6/1K6/8/3R4 w - - 0 1",
		"209": "1r6/8/8/5k2/8/1P6/1K6/4R3 w - - 0 1",
		"210": "8/k7/8/3P4/7r/2K5/1R6/8 w - - 0 1",
		"211": "3K4/3P2k1/8/8/8/8/2r5/5R2 w - - 0 1",
		"212": "6k1/r2KP3/8/8/8/8/8/5R2 w - - 0 1",
		"213": "8/1r2K3/4P1k1/8/8/8/8/R7 w - - 0 1",
		"214": "8/8/7k/4P3/r7/4K3/8/6R1 w - - 0 1",
		"215": "4r3/8/7k/8/8/4P3/4K3/6R1 w - - 0 1",
		"216": "1K6/1P1k4/1P6/8/8/r7/2R5/8 w - - 0 1",
		"217": "8/8/P7/8/3K4/r7/P2R4/k7 w - - 0 1",
		"218": "2R5/2P5/1k6/8/7r/1K5P/1P6/8 w - - 0 1",
		"219": "8/8/P5rk/8/8/KpR5/8/8 w - - 0 1",
		"220": "8/8/6P1/8/5p1k/r7/6K1/2R5 w - - 0 1",
		"221": "6k1/4K3/8/r4p2/4P3/8/8/5R2 w - - 0 1",
		"222": "8/8/2P2K2/8/6pr/6k1/R7/8 w - - 0 1",
		"223": "4k3/6K1/1P6/8/1p6/8/r7/6R1 w - - 0 1",
		"224": "3r4/P4Kp1/k7/8/8/8/1R6/8 w - - 0 1",
		"225": "2K5/2P2R2/k7/8/8/8/2r2p2/8 w - - 0 1",
		"226": "R6K/7P/p6k/r7/8/8/8/8 w - - 0 1",
		"227": "8/8/1P6/8/2k3pr/8/3RK3/8 w - - 0 1",
		"228": "8/6pP/6R1/7r/8/8/5K2/7k w - - 0 1",
		"229": "8/5R2/4PK2/2r5/8/2k1p3/8/8 w - - 0 1",
		"230": "8/8/2R3P1/8/7p/4k3/3r4/7K w - - 0 1",
		"231": "8/8/1R4P1/5p2/3K4/k3p3/4r3/8 w - - 0 1",
		"232": "K7/2rp4/8/2p3P1/8/8/5R2/7k w - - 0 1",
		"233": "8/R1p5/5pP1/8/8/7K/1k6/2r5 w - - 0 1",
		"234": "7k/8/4K1Rp/6pP/5p1r/8/8/8 w - - 0 1",
		"235": "k7/p2K1p2/7R/r1p1P3/1p6/8/8/8 w - - 0 1",
		"236": "8/2k5/2P3Pr/8/8/8/6p1/2R3K1 w - - 0 1",
		"237": "1R6/2pk4/1P6/8/3r4/K7/P7/8 w - - 0 1",
		"238": "2r5/6k1/6p1/P7/3P4/8/6K1/3R4 w - - 0 1",
		"239": "6R1/3r3p/2p2pP1/8/p1P5/8/5K1k/8 w - - ",
		"240": "5R2/1n6/P6k/8/1P5r/8/3K1p2/8 w - - 0 1",
		"241": "5r1k/R5p1/6Pp/5P1K/7P/1p1p4/8/8 w - - ",
		"242": "8/r4kp1/5p1p/3R1K1P/4PPP1/8/8/8 w - - ",
		"243": "8/Q5K1/8/8/8/8/3kp3/8 w - - 0 1",
		"244": "8/8/8/4K3/8/8/3Q1p2/6k1 w - - 0 1",
		"245": "8/8/8/8/6K1/6Q1/2p5/3k4 w - - 0 1",
		"246": "5Q2/8/8/8/K7/8/2p5/1k6 w - - 0 1",
		"247": "Q7/8/8/8/8/8/pk2K3/8 w - - 0 1",
		"248": "Q7/8/8/3K4/8/8/pk6/8 w - - 0 1",
		"249": "8/KQ6/8/8/8/8/p7/k7 w - - 0 1",
		"250": "7Q/8/1K6/8/p7/8/7p/6k1 w - - 0 1",
		"251": "3Q4/kr6/2K5/8/8/8/8/8 w - - 0 1",
		"252": "7Q/8/8/8/8/8/6rp/4K2k w - - 0 1",
		"253": "8/P7/R7/8/2q5/5K2/7k/8 w - - 0 1",
		"254": "8/5P1k/5Q2/7q/8/6K1/8/8 w - - 0 1",
		"255": "8/3KP1q1/8/8/8/4Q3/k7/8 w - - 0 1",
		"256": "6K1/5P2/6Q1/3q4/1k6/8/8/8 w - - 0 1",
		"257": "K7/1P6/k1q5/8/8/1Q6/8/8 w - - 0 1",
		"258": "6k1/5q2/3Pp2K/8/4Q3/8/8/8 w - - 0 1",
		"259": "8/7q/2K2p2/4p3/2k1P2p/8/3P4/7Q w - - ",
		"260": "2b4q/4K3/4p3/p2p1pk1/8/1P3P2/2P5/1Q6 w ",
		"261": "8/8/8/8/5k2/3q1N2/Q4K2/8 w - - 0 1",
		"262": "8/5B2/q7/8/4k3/8/5K2/3Q4 w - - 0 1",
		"263": "1QK3kq/6p1/8/6B1/8/8/8/8 w - - 0 1",
		"264": "8/8/8/8/1q6/3Q3K/p7/k1N5 w - - 0 1",
		"265": "8/2K5/2N5/pk6/q7/4Q3/8/8 w - - 0 1",
		"266": "K7/8/8/5q2/3Q4/k6p/4N3/8 w - - 0 1",
		"267": "8/8/1q1p4/1ppp4/8/5N1K/5k2/3Q4 w - - ",
		"268": "1R6/8/8/6K1/8/2Q5/qr6/k7 w - - 0 1",
		"269": "3q3b/2pp4/8/2B5/8/1k6/3P3Q/3K4 w - - ",
		"270": "3q2r1/6P1/p7/8/8/k3N1Qp/8/2K5 w - - 0 1",
		"271": "8/8/2Q4K/5p2/1k1qn1N1/1P3p2/5P2/8 w - ",
		"272": "6q1/7k/1r2NQ2/1p4p1/5PP1/1K3p2/8/8 w - ",
		"273": "8/K6N/8/2N5/1n6/6Q1/6pn/7k w - - 0 1",
		"274": "3KNB1k/7p/8/8/8/8/3n4/8 w - - 0 1",
		"275": "3K4/1p5p/1p6/8/8/1B6/1p6/b1k3B1 w - - ",
		"276": "8/8/4P3/k1N4r/8/K3N3/8/8 w - - 0 1",
		"277": "8/5B2/p7/q5B1/k7/3K4/1P6/8 w - - 0 1",
		"278": "2b4k/8/5Pr1/5N2/8/8/8/K1B5 w - - 0 1",
		"279": "N7/1p1kqP2/8/3P4/KP6/4B3/8/8 w - - 0 1",
		"280": "8/8/8/8/5N2/p1RK4/pk6/8 w - - 0 1",
		"281": "2k5/6R1/N7/8/8/6K1/7p/5b2 w - - 0 1",
		"282": "8/2k3Nr/8/3K4/1R6/8/8/8 w - - 0 1",
		"283": "3Nk3/r7/7R/8/3K4/8/8/8 w - - 0 1",
		"284": "R7/8/8/8/4K2p/8/8/3k1N1r w - - 0 1",
		"285": "8/b7/8/8/7K/1R6/2n5/N2k4 w - - 0 1",
		"286": "8/8/8/1K2k1n1/1n6/5R2/8/4N3 w - - 0 1",
		"287": "8/8/7n/3k4/B4p2/8/8/b2K2R1 w - - 0 1",
		"288": "7k/7p/B7/1K6/8/1n3b2/8/5R2 w - - 0 1",
		"289": "8/8/8/5p2/8/3R4/P7/1k1K1B1r w - - 0 1",
		"290": "3N4/4R3/6k1/8/5K1P/8/5p2/8 w - - 0 1",
		"291": "q3N1R1/p7/8/8/4P2k/8/4K2P/8 w - - 0 1",
		"292": "3N4/2p5/8/3q4/3k4/8/3PKP2/R7 w - - 0 1",
		"293": "5q2/2R5/5p2/1k3N2/2p1P3/2P5/2K5/8 w - ",
		"294": "1q3k2/7K/1P1P1p2/R2P4/8/B5b1/8/8 w - - ",
		"295": "2K5/4p3/1p1p1p2/1p2k3/1P3q2/8/1PBPPP2/7",
		"296": "RK6/8/1k6/7R/8/8/pp6/8 w - - 0 1",
		"297": "8/8/8/8/7k/6pq/R1R5/4K3 w - - 0 1",
		"298": "8/4r3/1P5K/8/1p2kr2/8/R7/R7 w - - 0 1",
		"299": "5kr1/7R/3K4/7R/2p5/2P1p1r1/8/8 w - - ",
		"300": "1B6/8/5kP1/8/8/7K/8/1r1B2N1 w - - 0 1",
	}

	rx := regexp.MustCompile(`(?s)([0-9]+)VAR-(\d+)\n(.+?)\n\n`)

	matches := rx.FindAllStringSubmatch(text, -1)
	lastNumber := "-1"
	for _, match := range matches {
		moves := []string{}
		comment := []string{}
		for _, line := range strings.Split(match[3], "\n") {
			if ok, _ := regexp.Match(`^([0-9]+).+$`, []byte(line)); ok {
				moves = append(moves, line)
			} else {
				comment = append(comment, line)
			}
		}
		movesString := strings.Join(moves, "\\n")
		commentString := strings.Join(comment, " ")
		if lastNumber != match[1] {
			fmt.Println(match[1])
			lastNumber = match[1]
		}
		fmt.Printf("{ initialBoard: '%v',\nactions: '%v',\ncomment: '%v'\n},\n", initialBoards[lastNumber], movesString, commentString)
	}
}
