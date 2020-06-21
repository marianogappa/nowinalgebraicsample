package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	text := `2VAR-1
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
and Black 
gets a draw.

19VAR-1
1 b5    h4
2 b6    h3
3 b7    h2
4 b8=Q+
with Pawn loss
Black loses.

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

26VAR-1
1 a4        Ke4
2 a5        Kd5
and Black catches 
the Pawn. 

26VAR-2
1 Kf5!     c5
2 Ke5    Ke3
3 Kd5    Kd3
4 Kxc5    
wins for White.

26VAR-3
	1 Kf5!    Ke3
2 Ke5    c6 
3 Kd6    Kd4
4 Kxc6    Kc4
	and Black draws.

26VAR-4
1 Kf5!    Ke3
2 Ke5    Kd3
3 Kd5    c6+
4 Kc5    Kc3
5 a4
and White wins.

26VAR-5
1 Kf5!    Ke3
2 Ke5    c6
3 a4         Kd3
4 Kd6     Kc4 
and Black saves 
himself.

26VAR-6
1 Kf5!    Ke3
2 Ke5    c6
3 a4         Kd3
4 a5        c5
5 a6        c4
6 a7        c3
7 a8=Q    c2
8 Qe4+    Kd2
9 Qd4+    Ke2
10 Qc3    Kd1
11 Qd3+    Kc1
12 Kd4    Kb2
13 Qe2      Kb1
14 Kc3    c1=Q+
15 Kb3
and Black will be 
mated. 

26VAR-7
1 Kf5!    Ke3
2 Ke5    c6
3 a4         Kd3
4 a5        c5
5 a6        c4
6 a7        c3
7 a8=Q    c2
8 Qe4+    Kd2
9 Qd4+    Ke2
10 Qc3    Kd1
11 Qd3+    Kc1
12 Kd4    Kb2
13 Qe2      Ka1!
14 Qxc2 
and Black is
stalemated.

26VAR-8
1 Kf5!    Ke3
2 Ke5    c6
3 a4         Kd3
4 a5        c5
5 a6        c4
6 a7        c3
7 a8=Q    c2
8 Qd5+!    Kc3
9 Qd4+    Kb3
10 Qa1    Kc4
11 Ke4
Black Pawn falls.

26VAR-9
1 Kf5!    Ke3
2 Ke5    c6
3 a4         Kd3
4 a5        c5
5 a6        c4
6 a7        c3
7 a8=Q    c2
8 Qd5+!    Ke3
9 Qg2    c1=Q
10 Qg5+
winning the Queen.

26VAR-10
	1 Kf5!    Ke3
2 Ke5    c6
3 a4         Kd3
4 a5        c5
5 a6        c4
6 a7        c3
7 a8=Q    c2
8 Qd5+!    Ke3
9 Qg2    Kd3
10 Qg5     Kc3
11 Qc1
	is decisive. 

27VAR-1
1 Kf7    g5
2 Kg7    Kb3
3 Kxh7    Kc4
4 Kg6    g4
5 Kf5    Kd5
6 Kxg4    Ke6
7 Kg5    Kf7
	and an automatic
	draw.

27VAR-2
1 h4!    h6
2 h5    Kb3
3 Kf7
White wins at once. 

27VAR-3
1 h4!    Kb3
2 Kf7    Kc4
3 Kxg7    h5
4 Kg6    Kd5
5 Kxh5    Ke6
6 Kg6    Ke7
7 Kg7
Black King kept at 
arm's length and wins.

27VAR-4
1 h4!    h5
2 Kf7    g5
3 hxg5    h4
4 g6    h3
	both sides Queen
	with a drawn result.

27VAR-5
1 h4!    h5
2 Kf8!    g5
3 hxg5
White's Pawn will 
Queen with check.

27VAR-6
1 h4!    h5
2 Kf8!    g6
3 Kf7      g5
4 hxg5    h4
	leads to a draw.

27VAR-7
1 h4!    h5
2 Kf8!    g6
3 Kg7      g5
4 hxg5    h4
	leads to a draw.

27VAR-8
1 h4!    h5
2 Kf8!    g6
3 Ke7!    Kb3
4 Kf6
White wins with 
both Pawns 
captured. 

28VAR-1
1 Kf5    Kf7
2 Ke5    Ke7
3 Kd5    Kd7
4 h6
	makes Black regret it.

28VAR-2
1 Kf5    Kh6
2 Ke5    Kxh5
3 Kd5    Kg6
4 Kc6    Kf6
5 Kxb6    Ke7
6 Kc7
	White wins.

28VAR-3
1 h6+    Kh7
2 Kh5     Kh8
3 Kg6     Kg8
4 h7+     Kh8
5 Kh6     b5
6 a5         b4
7 a6         b3
8 a7         b2
White wins with 
9 a8=Q#.

29VAR-1 
1 f5        Kf7
2 Ke5    Ke7
3 f6+    Kf7
4 Kf5    Kf8
5 Ke6    Ke8
6 f7+    Kf8
7 Kf6 
draws by
stalemate

29VAR-2
1 f5        Kf7
2 Ke5    Ke7
3 f6+    Kf7
4 Kf5    Kf8
5 Ke6    Ke8
6 f7+    Kf8
7 Kd6    Kxf7
8 Kc6    Ke7
9 Kb6    Kd7
10 Kxa5    Kc7
11 Kb5    Kb8
Black draws against 
the Rook Pawn.

29VAR-3
1 Kd5    Kf5
2 Kc5    Kxf4
3 Kb5    Ke5
4 Kxa5    Kd6
5Kb6     Kd7 
6 Kb7    Kd8
7 a5
White wins

30VAR-1
1 Kf5    Kd6
2 Kf6    Kc5
3 Ke6    Kxb5
4 Kd5    Kb6
White loses 
both Pawns. 

30VAR-2
1 Kd5!    Kf7
2 b6     
White wins 
instantly.

30VAR-3
1 Kd5!    Kd8
2 Ke6    Ke8 
3 b6        Kd8
4 b7
White mates 
next move.  

30VAR-4
1 Kd5!     Kd8
2 Ke6    Ke8
3 b6        cxb6
4 c7 
White mates 
next move. 

31VAR-1
1 Kc5    b6+
2 axb6+    Kb7
Black draws

31VAR-2
1 Kc5    b6+
2 Kc6     bxa5
3 Kc7    a4
4 b6+    Ka6
5 b7    a3
6 b8=Q
White wins Pawn
and game.

31VAR-3
1 Kc5    Ka8
2 Kb6    Kb8 
3 a6        Ka8
4 axb7+    Kb8
5 Ka6    Kc7
6 Ka7
White Pawn 
Queens and wins.

31VAR-4
1 Kc5    Ka8
2 Kb6    Kb8
3 a6        bxa6
4 Kxa6!    Ka8
5 b6
and White wins.

32VAR-1
1 f6+    Kf7
2 Kf5    Kf8
3 Ke6
It’s all over with 
lost Pawn.
	
	32VAR-2
1 f6+    Kf8!
2   f7!        Ke7
3 f8=Q+     Kxf8
4 Kf6    Ke8
5 Ke6    Kd8
6 Kxd6    Ke8
7 Kc7
	White wins.

32VAR-3
			1 f6+    Kf8!
2 f7!        Kxf7
3 Kf5        Ke7
4 Kg6    Kd8
5 Kf6    Kd7
6 Kf7        Kd8
7 Ke6    Kc7
8 Ke7    Kc8
9 Kx d6     Kd8
10 Kc6        Ke7
11 Kc7
and the Pawn has 
a clear road.

32VAR-4
1 f6+    Kf8!
2 f7!        Kxf7
3 Kf5        Ke7
4 Kg6    Kd8
5 Kf6    Kd7
6 Kf7        Kd8
7 Ke6    Kc7
8 Ke7    Kc8
9 Kxd6     Kd8
10 Kc6    Kc8
11 d6        Kd8
12 d7        Ke7
13 Kc7
and wins

33VAR-1
1 Kf6      Ke8
2 Ke6            Kd8
3 d7            Kc7
4 Ke7            
White wins

34VAR-1
1 d5        cxd5
2 Kxd5    Kb6
3 Kc4    Kc6
4 b5+    Kb6
	the position is a 
familiar King on 
King draw.

35VAR-1
1 Kd6    Kd8
2 Ke6    Ke8
3 c6        bxc6
4 b7
White Queen’s
first and wins.    

35VAR-2
1 Kd6    Kd8
2 Ke6    Ke8
3 c6        Kd8
4 cxb7    Ke8
5 b8=Q#
Clean position.

36VAR-1
1 Kd7!    Ka7
2 Kc7    Ka6
3 Kb8    Ka5
4 Kb7    Kxa4
5 Kxb6     Kb4
6 c5
seizing the opposition
White gets behind the
Pawn and wins.  

36VAR-2
1 Kd7!    Kb7
2 a5!        Ka6
3 axb6    Kxb6
4 Kd6    Kb7
5 c5        Kc8
6 Kc6 
White takes familiar 
winning position.

37VAR-1
1 d6        Kc8
2 b6        cxb6
3 Kxb6    Kd7
drawing easily. 

37VAR-2
1 d6        Kc8
2 d7+    Kd8
3 Kb7    Kxd7
4 Kb8    Kd6
	and Black draws.

37VAR-3
1 Kb7    Kd7
2 Kb8    Kd6
3 Kc8    Kxd5
4 Kxc7
and White wins.

38VAR-1
1 Kf5    Kf8
2 Kg6          
White teases his
only Pawn up the 
g file to Queen.

38VAR-2
1 Kf5    Kf7
2 Kg5    Kf8
3 Kxh5    
etc.

39VAR-1
1 g4        Kh8
2 g5    Kg8
3 g6    hxg6
4 Kxg6    Kh8
Black draws against 
the Rook Pawn

39VAR-2
1 g4     Kh8
2 g5    Kg8
3 g6    hxg6
4 hxg6    Kh8
5 g7+
	with 5 . . .Kg8
	a draw occurs.

39VAR-3
1 g3        Kh8
2 g4             Kg8
3 g5        Kh8
4 g6        Kg8
5 g7    Kf7
6 Kxh7
	White wins.

40VAR-1
1 Kf2    Ke5
2 g6    Kf6
3 h5    Ke6
4 g7        Kf7
5 h6
Black's King cannot
be all over the board 
at the same time.

41VAR-1
1 Kd5    Kc8
2 Kd4!    Kc7
3 Kc5    Kb8
4 Kb6 
Black's Pawn is lost.

41VAR-2
1 Kd5    Kc8
2 Kd4!    Kd8
3 Kc4    Kc8
4 Kd5    Kd8
5 Kd6    Kc8
6 c7        Kb7
7 Kd7    Ka7
8 c8=Q
stalemate

41VAR-3
1 Kc4    Kc8
2 Kd5    Kd8
3 Kd6    Kc8
4 c7        Kb7
5 Kd7    Ka7
6 Kc6     Ka8
7 c8=Q+     Ka7
and mate next 
with 8 Qb7#.

42VAR-1
1 c5        b6
2 c6    Kc8
3 Kxb6 
and it's all over.

42VAR-2
1 c5         Kc8
2 Kb6    Kb8
3 c6         Ka8
4 c7 
resulting in a stalemate.

42VAR-3
1 c5        Kc8
2 Kb6    Kb8
3 c6        Kc8
4 c7        Kd7
5 Kxb7
White wins

42VAR-4
1 c5        Kc8
2 Kb6    Kb8
3 c6        bxc6
4 Kxc6    Kc8
5 Kb6    Kb8
6 b5        Kc8
7 Ka7
White wins.

42VAR-5
1 c5        Kc8
2 Kb6    Kb8
3 c6        Ka8
4 cxb7+    Kb8
5 Ka6
White wins

42VAR-6
1 c5        Kc8
2 Kb6    Kb8
3 c6        Ka8
4 Kc7    bxc6
5 Kxc6    Ka7
6 b5        Kb8
7 b6        Kc8
Black escapes 
with a draw.

43VAR-1
1 f3+    gxf3
2 gxf3+    Kf4
Black has a draw.

43VAR-2
1 g3        Ke5
2 Ke3    Kf5
3 Kd4    Kf6
4 Ke4    Kg5
5 Ke5    Kg6
6 Kf4    Kh5
7 Kf5    Kh6
8 Kxg4
	White wins.

43VAR-3
1 g3         Kd4
2 f3        Ke5
3 fxg4    Ke4
4 Kf2    Ke5
5 Ke3    Ke6
6 Kf4    Kf6
7 g5+    Kg6
8 Kg4    Kg7
9 Kf5    Kf7
10 g6+    Kg7
11 Kg5    Kg8
12 Kf6    Kf8
13 g7+    Kg8
14 g4    Kh7
15 Kf7
	White wins

43VAR-4
1 g3        Kd4
2 f3        gxf3+
3 Kxf3    Ke5
4 Kg4!    Kf6
5 Kh5    Kg7
6 Kg5    Kf7
7 Kh6     
etc. all the way to
a Queen.
	
	44VAR-1
1 Kb4    c5+
2 dxc5+    Kc6
3 Kb3    Kxc5 
and the position 
is a draw. 

44VAR-2
1 Kb3    Kc7
2 Kc3    Kd6
3 Kd3    Kd7
4 Ke4     Ke6
5 c5        Kf7
6 Kf5    Ke7
7 Ke5    Kd7
8 Kf6    Kd8
9 Ke6    Kc7
10 Ke7    Kc8
11 Kd6    Kb7
12 Kd7
	Black's Pawn falls.

44VAR-3
1 Kb3    Kc7
2 Kc3    Kd6
3 Kd3    Kd7
4 Ke4     Ke6
5 c5        Kf6
6 d5!    cxd5
7 Kxd5    Ke7
8 Kc6    Kd8
9 Kb7 
assures White
Pawn's future.

44VAR-4
1 Kb3    Kc7
2 Kc3    Kd6
3 Kd3    Kd7
4 Ke4     Ke6
5 c5        Kf6
6 d5!    Ke7
7 5xc6    Kd8
8 Ke5    Kc7
9 Kd5    Kc8
10 Kd6    Kd8
11 c7+    Kc8 
and Black has 
a draw.

44VAR-5
1 Kb3    Kc7
2 Kc3    Kd6
3 Kd3    Kd7
4 Ke4     Ke6
5 c5        Kf6
6 d5!        Ke7
7 d6+    Kd7
8 Ke5    Kd8
9 d7!    Ke7
10 d8=Q+    exd8
11 Kd6    Kc8
12 Kxc6
leaving a standard 
winning position.

45VAR-1
1 Kf1    Kd4
2 Kf2    Kc5
3 e4!    Kb6
4 e5        Kxa6
5 e6 
White's Pawn cannot 
be caught and wins.

45VAR-2
1 Kf1    Kd4
2 Kf2    Kc5
3 e4!    Kd4
4 Kf3    Ke5
5 Ke3    Ke6
6 Kd4    Kd6
7 e5+    Ke6
8 Ke4    Ke7
9 Kd5    Kd7
10 e6+    Ke7
11 Kc6     Kxe6
12 Kb7    Kd7
13 Kxa7     Kc7
14 Ka8    Kc8
White cannot 
extricate his King, 
hence a crafty draw.

45VAR-3
1 Kf1    Kd4
2 Kf2    Kc5
3 e4!    Kd4
4 Kf3    Ke5
5 Ke3    Ke6
6 Kd4    Kd6
7 e5+    Ke6
8 Ke4    Ke7
9 Kd5    Kd7
10 e6+    Ke7
11 Kc6    Ke8
12 Kd6    Kd8
13 Kc6    Kc8
14 e7     Kb8
mates next move 
with 15 e8=Q#.

46VAR-1
1 Kc6    Ke7
2 Kd5    f6
3 e6    f5
4 Ke5
wins the Pawn.

46VAR-2
1 Kc6    Ke7
2 Kd5    f6
3 e6        Ke8
4 Kd6    Kd8
5 e7+    Ke8
6 g4        Kf7
7 Kd7
	White wins.

46VAR-3
1Kc6    Ke7
2 Kd5    Kd7
3 Ke4    Ke6
4 Kf4    f6
5 exf6    Kxf6
6 Kg4    Kg6
Black has the 
opposition and 
draws.

46VAR-4
1 Kc6    Ke7
2 Kd5    Kd7
3 Kd4!    Ke7
4 Ke3    Kd7
5 Kf4    Ke6
6 Ke4    f6
7 exf6     Kxf6
8 Kf4    Kg6
9 Kg4     Kf6
10 Kh5
	White has the opposition 
and wins.

46VAR-5
1 Kc6    Ke7
2 Kd5     Kd7
3 Kd4!    Ke7
4 Ke3     Kd7
5 Kf4    Ke6
6 Ke4     Kd7
7 Kf5     Ke7
8 g4     Ke8!
9 e6     Ke7
10 exf7     Kxf7
11 Kg5     Kg7
again Black takes the
opposition and draws.

46VAR-6
1 Kc6    Ke7
2 Kd5    Kd7
3 Kd4!    Ke7
4 Ke3    Kd7
5 Kf4    Ke6
6 Ke4    Kd7
7 Kf5    Ke7
8 g4        Ke8!
9 Kf6!    Kf8
10 e6     fxe6
11 Kxe6    Kg7
12 Kf5    Kf7 
lets Black draw
	easily.

46VAR-7
1Kc6    Ke7
2 Kd5    Kd7
3 Kd4!    Ke7
4 Ke3    Kd7
5 Kf4    Ke6
6 Ke4    Kd7
7 Kf5    Ke7
8 g4        Ke8!
9 Kf6!    Kf8
10 g5    Ke8
11 e6    Kf8
12 e7+    Ke8
and White must 
concede the draw.

46VAR-8
1Kc6    Ke7
2 Kd5    Kd7
3 Kd4!    Ke7
4 Ke3    Kd7
5 Kf4    Ke6
6 Ke4    Kd7
7 Kf5    Ke7
8 g4        Ke8!
9 Kf6!    Kf8
10 g5    Ke8
11 Kg7!    Ke7
12 Kg8    Ke6
13 Kf8    Kxe5
14 Kxf7
and White wins. 

47VAR-1  
1 Kg4     Kh2
2 Kg5    Kg3
3 Kg6    Kxf4
	leads to a draw.

48VAR-1
1 Kb5    Kc3 
2 Kc5    Kd3 
3 Kd5    Ke3
4 Ke5    Kf3
5 Kf5    Kg3
6 Kg6     Kxh4
	and a draw.

49VAR-1
1 Kf5    Kf3
2 Kg6    Kg3
3 h5        Kh4
4 Kxg7     Kxh5
	with draw as result.

49VAR-2
1 Ke6!    Kf4
2 Kf7    g5
3 h5        Kg3
4 Kg6
White wins

50VAR-1
1 e6!     Ke7
2 exf7    Kxf7
3 Kd5    Ke7
4 Ke5 
again White 
wins with the 
opposition

50VAR-2 
1 e6!    f6
2 Kc5    Ke7
3 Kd5    Ke8
4 Kd6    Kd8
5 e7+    Ke8
6 Ke6    f5
7 Kxf5    Kxe7
8 Ke5
White opposition 
and extra Pawn’s 
deflection wins.

52VAR-1
1 Kf1    Kc3
2 Ke2    Kd4
3 g4        Ke4
4 g3        Kd4
5 g5        Ke5
6 g4        Ke6
7 Kxe3    Kf7
8 Kf4    Kg6
Pawn loss gives 
Black an easy draw.

52VAR-2
1 Kf1    Kc3
2 Ke2    Kd4
3 g4        Ke4
4 g3        Kd4
5 g5        Ke5
6 g4        Ke6
7 Kxe3    Kf7
8 Ke4    Kg6
9 Kf4    Kf7
10 Kf5    Kg7
11 g6        Kh6
12 Kf6
clumsy stalemating 
of Black.

52VAR-3
1 Kf1    Kc3
2 Ke2    Kd4
3 g4        Ke4
4 g3        Kd4
5 g5        Ke5
6 g4        Ke6
7 Kxe3    Kf7
8 Ke4    Kg6
9 Kf4    Kf7
10 Kf5    Kg7
11 g6        Kh6
12 g7     Kxg7
13 Kg5
White having the
opposition wins.

52VAR-4
1 Kf1    Kc3
2 Ke2    Kd4
3 g4        Ke4
4 g3        Kd4
5 g5        Ke5
6 g4        Ke6
7 Kxe3    Kf7
8 Ke4    Kg6
9 Kf4    Kf7
10 Kf5    Kg7
11 g6        Kg8
12 Kf6    Kf8
13 g7+    Kg8
14 g5        Kh7
15 Kf7
stalemating Black,
deprives White 
of the win.

53VAR-1
1 Kg5    Kf3
2 h4        Kg2
3 h5        Kh3!
4 Kg6    Kh4
5 h3    Kxh3
6 Kxg7
White wins.

53VAR-2
1 Kg5    Kf3
2 h4        Kg2
3 h5        Kh3!
4 Kg6    Kg4!
5 h3+    Kh4
6 Kxg7    Kxh5
Black draws. 

53VAR-3
1 Kg4    Kf2
2 h4        Kg2
3 h3        g6
4 Kg5    Kxh3
and the position 
is drawn.

53VAR-4
1 Kg3!    Kf1 
2 h4    Ke2
3 h5    Ke3
4 h4    Ke4
5 Kg4    Ke3
6 Kf5    Kf3
7 h6        gxh6
8 h5
White wins

53VAR-5
1 Kg3!    Kf1
2 h4    g6
3 Kf4!    Kg2
4 h5!    gxh5
5 h4
White wins again.

53VAR-6
1 Kg3!    Ke3 
2 h4        Ke4
3 Kg4    Ke5
4 Kg5    Ke4
5 h5        Kf3
6 Kg6    Kg4
7 h3+    Kh4
ends in a draw.

54VAR-1
1 a5        Kc6
2 a6        Kb6
3 a7        Kxa7
4 Kxc5    Kb7
5 Kd6      Kc8
6 Ke6     Kd8
7 Kxf5    Ke7
8 Kg6    Kf8
9 Kf6 
White wins with 
or without move.

55VAR-1
1 a4        Kf7
2 a5        Ke7
3 b6        Kd7
4 bxa7    Kc7
White Queens 
and wins.

55VAR-2
1 a4        Kf7
2 a5        Ke7
3 b6     axb6
4 a6!
again White wins.

55VAR-3
1 a4        h5
2 a5        h4
3 b6     axb6
4 axb6    h3
5 b7        h2 
6 b8=Q    h1=Q
position is a draw.

56VAR-1
1 Kf4    h5
2 Kg5    d5
3 Kxh5    d4
4 Kg4    d3
5 Kf3
White overtakes 
the Pawn.

56VAR-2
1 Kf4    h5
2 Kg5    Kb6
3 Kxh5    Kc7
4 Kg5    Kb6
5 Kf5    Kc7
6 Ke6
and the second 
Pawn goes.

56VAR-3
1 Kf4    Kb6
2 Kf5    Kc7
3 Kf6    h5
4 Kg5
White in time 
to catch the 
second Pawn.

56VAR-4
1 Kf4    Kb6
2 Kf5    Kc7
3 Kf6    Kb6
4 Ke6    h5
5 Kxd6    h4
6 c7    h3
7 c8=Q    h2
8 Qa6#

57VAR-1
1 Kf7!    Kh8
2 Kg6
wins both Black 
Pawns.

57VAR-2
1 Kf7!    h5
2 Kf6    hxg4
3 hxg4    Kh8
4 Kxg5    Kg7
Black gets the 
opposition with
easy draw.

57VAR-3
1 Kf7!    h5
2 h4!    hxg4
3 hxg5    g3
4 g6+    Kh6
5 g7        g2
6 g8=Q
White wins.

57VAR-3
1 Kf7!    h5
2 h4!    gxh4
3 g5    h3
4 g6+    Kh6
5 g7    h2
6 g8=Q    h1=Q
7 Qg6#

57VAR-3
1 Kf7!    h5
2 h4!    Kh6
3 Kf6!    gxh4
4 g5+    Kh7
5 Kf7
and White wins.

58VAR-1
1 Kf3    Kg8
2 Ke3    Kf8
3 Kd3    Ke7
4 Kc3    Kf6
5 Kb3    Kxg6
6 Kxa3    Kf5
7 Kb4    g5
8 a4
both sides Queen. 
and draw.

58VAR-2
1 Kf4!    Kg8
2 Ke5    Kf8
3 Kd6!    Ke8
4 Ke6    Kd8
5 Kf7
Blacks’ outreach 
is fatal.

58VAR-3
1 Kf4!    Kg8
2 Ke5    Kf8
3 Kd6!    Ke8
4 Ke6    Kf8
5 Kd7    Kg8
6 Ke7    Kh8
7 Kd6    Kg8
8 Kc5    Kf8
9 Kb4    Ke7
10 Kxa3    Kf6
11 Kb4    Kxg6
12 a4        Kf5
13 a5        g5
14 a6    g4
15 a7    g3
16 a8=Q    Kf4
17 Qg2
and White wins.

59VAR-1
1 Kf2    g3+
2 Kg2
stalemate

59VAR-2
1 Kf2    gxf2
2 Kxf3    Kh3
3 f5        h4
4 f6,        Kh2
5 f7        h3
6 f8=Q    Kg1
7 Qh8    h2
8 Qa1#

59VAR-3
1 Kf2    Kh3
2 f5        Kh2!
Black will be 
first to Queen.

59VAR-4
1 Kf2    Kh3
2 Kg1    Kh4
3 Kg2    g3
4 Kh1    Kh3
5 f5        g2+
6 Kg1    Kg3
7 f6        h4
8 f7        h3
9 f8=Q    h2#!
serving White a 
bewildering blow!

59VAR-5
1 Kf2    Kh3
2 Kg1    Kh4
3 Kg2    g3
4 Kg1    g2
5 Kxg2
stalemating Black
is unthinkable.

60VAR-1
1 Kc4!    a4
2 d5    a3
3 Kb3
hopeless for Black.

60VAR-2
1 Kc4!    Kxf4
2 d5        Ke5
3 Kc5    f4
4 d6        Ke6
5 Kc6    f3
6 d7    f2
7 d8=Q    f1=Q
8 Qe8+    Kf6
9 Qf8+
White wins the 
Queen.

61VAR-1
1 c5        Ke6
2 Kd4    a4
3 Kc4    a3
4 Kb3    Kd5
White loses 
instead of winning.

61VAR-2
1 Kd4!    a4
2 Kc3    Ke6
3 Kb4
wins for White.

61VAR-3
1 Kd4!    Ke6
2 Kc5    Ke5
3 Kb5    Kd4
4 c5        a4
5 c6        a3
with a comfortable 
draw.

61VAR-4
1 Kd4!    Ke6
2 Kc5    Ke5
3 Kb5    Kd4
4 g5!    fxg5
5 c5        g4
6 c6        g3
7 c7        g2
8 c8=Q    g1=Q
9 Qc5
and White wins 
the Queen.

62VAR-1
1 Kd8    Kc6
2 Ke7    f5!
3 gxf5    g4
both sides Queen 
with draw result.

62VAR-2
1 a5!    Kc6
2 Kb8!    Kb5
3 Kb7    f5
4 a6        fxg4
5 a7        g3
6 a8=Q    g2
7 Qa1
White wins.

62VAR-3
1 a5!    Kc6
2 Kb8!    Kb5
3 Kb7    Kxa5
4 Kc6    Ka6
5 Kd5    f5
6 gxf5    g4
7 Ke4
Black's Pawn 
is lost.

62VAR-4
1 a5!    Kc6
2 Kb8!    Kb5
3 Kb7    Kxa5
4 Kc6    Kb4
5 Kd5    f5
6 gxf5    g4
7f6
Pawn reaches f8, 
with check win.

63VAR-1
1 Ke5    h5
2 c5     h4
3 c6    h3
4 c7        h2
5 c8=Q    h1=Q
6 Qc4+    Ka3
7 Qb3#

63VAR-2
1 Ke5    h5
2 c5        Kb5!
3 Kd6    h4
4 c6        h3
5 c7        h2
6 c8=Q    h1=Q
7 Qc5+     Ka6!
results in a draw.

64VAR-1
1 b6        b3
2 b7        b2
3 b8=Q    b1=Q
4 Qxb1
with stalemate.

65VAR-1
1 Ke4    Kd6
2 Kd4    Kc6
3 Ke5    Kb6
4 Kd5    Ka6
5 Kc5    Ka7
5 Kxb5
White wins.

65VAR-2
1 Ke4    Kd6
2 Kd4    Kc6
3 Ke5    Kb6
4 Kf5    Kc6
5 Kg5    Kd6
6 Kxh5
White wins.

65VAR-3
1 Ke4    Kf6
2 Kd5    Kf5
3 Kc5    Kg4
4 Kxb5    Kxh4
5 Kc5    Kg4
6 b5        h4
7 b6        h3
8 b7        h2
9 b8=Q    h1=Q
and a draw.

66VAR-1
1 Ke7!    Kd2
2 Kd6    Kc3
3 Kc5    Kb3
4 Kxb5!
White wins.

66VAR-2
1 Ke7!    Kf2
2 Kf6    Kg3
3 Kg5    Kh3
4 Kxh5!
White wins again.

66VAR-3
1 Ke7!    Ke2
2 Ke6!    Kd3
3 Kd5    Kc3
4 Kc5    Kb3
5 Kxb5!
White wins.

66VAR-4
1 Ke7!    Ke2
2 Ke6!    Kf3
3 Kf5    Kg3
4 Kg5    Kh3 
5 Kxh5!    
White wins.

66VAR-5 
1 Ke7!    Ke2
2 Ke6!    Ke3
3 Ke5!    Ke2
4 Ke4!    Kd2
5 Kd4    Kc3
6 Kc5    Kxb3
7 Kxb5!
White wins

66VAR-6 
1 Ke7!    Ke2
2 Ke6!    Ke3
3 Ke5!    Ke2
4 Ke4!    Kf2
5 Kf4    Kg2
6 Kg5    Kh3
7 Kxh5!
White wins.

67VAR-0

68VAR-1 
1 a5        Kc5
2 Kb2!    Kb5
3 Kc3    Kxa5
4 Kd4    Kb5
5 Ke4    Kc6
6 Kf5    Kd7
7 Kxf6    Ke8
8 Kxg5    Kf7
9 Kh6    Kg8
10 Kg6
White wins.

68VAR-2 
1 a5        Kc5
2 Kb2!    Kb5
3 Kc3    Kxa5
4 Kd4    Kb5
5 Ke4    Kc6
6 Kf5    Kd7
7 Kxf6    Ke8
8 Kxg5    Kf7
9 Kh6    Kf6
10 g5+    Kf7
11 Kh7
and White wins.

68VAR-3 
1 a5        Kc5
2 Kb2!    f5
3 gxf5    g4
4 f6!    Kd6
5 a6        Ke6
6 a7        g3
7 a8=Q
White wins again.

69VAR-1
1 f4        d5
2 f5        d4
3 f6        d3
4 f7        d2
5 f8=Q+
White wins.

69VAR-2 
1 f4        Kb4
2 h4        Kc5
3 h5
Rook Pawn Queens 
and wins.

69VAR-3 
1 f4    Kb4
2 h4    d5
3 h5    d4
Black will Queen 
with check.

69VAR-4 
1 f4    Kb4
2 h4    d5 
3 f5!    d4
4 f6    
White will Queen
	with check.

69VAR-5 
1 f4    Kb4
2 h4    d5 
3 f5!    Kc5
4 h5    Kd6
5 h6    Ke7
6 h7    Kf6
7 h8=Q+
Queen routs all of
Blacks Pawns.

69VAR-6 
1 f4    Kb4
2 h4    d5 
3 f5!    Kc5
4 h5    d4
5 f6    Kd6
6 h6    Ke6
7 h7 
White wins.

70VAR-1 
1 a5        g5
2 hxg5    fxg5
3 Ke5
Black must still 
play to capture 
the passed Pawn.
but it’s hopeless.

71VAR-1
1 g4    a5
2 a4    Kf6
3 h4    Kg6
4 Kd5    Kf6
5 Kc5    Ke5
6 Kb5    Kf4
7 Kxa5    Kxg4
8 Kb6
	and White will
Queen the Pawn.

72VAR-1
1 g4+!    Kg6
2 g5
White works a passed 
Pawn on the King side.

72VAR-2
1 g4+!    hxg4+
2 Kg3    Ke5
3 Kxg4    Kd4
4 h5    Kc4
5 h6
Rook Pawn Queens 
long before Black 
can.  White wins.

73VAR-1
1 c6!    dxc6
2 d6    exd6
3 f5
Pawn marches 
merrily up the board 
and Queens.

74VAR-1
1 g6!    hxg6
2 hxg6    Kf8
3 Kd6    Ke8
4 Ke6    Kd8
5 Kf7
ends the struggle.

74VAR-2
1 g6!    hxg6
2 hxg6    Kf8
3 Kd6    Ke8
4 Ke6    Kf8
5 Kd7    Kg8
6 Ke7    Kh8
7 f6    gxf6
8 Kf7    f5
9 g7+ 
and White mates
in two.

74VAR-3
1 g6!    h6
2 f6+    gxf6+
3 Kf5    Kf8
4 Kxf6    Kg8
5 g7    Kh7
Black has a draw.

74VAR-4
1 g6!    h6
2 Kd5    Kd7
3 f6 
Pawn will Queen, 
and win.

74VAR-5
1 g6!    h6
2 Kd5    Kf6
3 Ke4    Kg5
4 Ke5    Kxh5
5 Ke6    Kg5
6 f6!    Kxg6
7 f7
8 f8=Q

75VAR-1
1 Kb4    Kd4
2 a4      Kd5
3 a5    Kd6
4 axb6
wins at once.

75VAR-2
1 Kb4    Kd4
2 a4      Kd5
3 a5    bxa5+
4 Kxa5    Kc5
5 Ka4    Kb6.
Black draws.

75VAR-3
1 Kb4    Kd4
2 a4    Kd5
3 a5    bxa5+
4 Ka4!    Kc5
5 Kxa5    Kd6
6 b6    Kc6
7 bxa7 
is conclusive.

76VAR-1
1 Ke2    Kb7
2 Kd3    Ka8
3 Kc4    f3
4 Kd3
winning the Pawn 
is the penalty.

77VAR-1
1 Ke4    Kxh2
2 g4
the passed Pawn 
can never be caught.

77VAR-2
1 Ke4    Kg4
2 h4    Kh5
3 Kf4    Kh6
4 g4    Kg6
5 h5+    Kh6
6 Ke4    Kg5
7 Kf3    Kh6
8 Kf4    Kh7
9 g5    Kg7
10 g6        Kh6
11 Kg4    d3
12 Kf3
will win the Pawn.

78VAR-1
1 a3!    Ka5
2 Kb7    f5
3 gxf5    g4
4 f6    g3
5 f7    g2
6 f8=Q    g1=Q
7 Qb4#

78VAR-2
1 a3!    Ka5
2 Kb7    Kxa4
3 Kc6    Kxa3
4 Kd5    Kb4
5 Ke5    Kc5
6 Kf6    Kd6
7 Kxf7    Ke5
8 Kg6     Kf4
9 Kh5    Ke5
10 Kxg5    Ke6
11 Kg6    Ke7
12 g5    Kf8
13 Kh7
White wins.

78VAR-3
1 a3!    f5
2 gxf5    g4
3 f6        g3
4 f7        g2
5 f8=Q    g1=Q
6 Qc8+    Kb6
7 Qb7+    Ka5
8 Qb5# 

78VAR-4
1 a3!    f5
2 gxf5    g4
3 f6        g3
4 f7        g2
5 f8=Q    g1=Q
6 Qc8+    Kb6
7 Qb7+    Kc5
8 Qa7+
White wins the 
Queen.

78VAR-5
1 a3!    f5
2 gxf5    g4
3 f6        g3
4 f7        g2
5 f8=Q    g1=Q
6 Qc8+    Ka5
7 Qc3+    Ka6
8 Qc4+    Kb6
9 a5+    Kxa5
10 Qb4+    Ka6
11 Qa4+    Kb6
12 Qa7+
again White wins
the Queen.

78VAR-6
1 a3!    f5
2 gxf5    g4
3 f6        g3
4 f7        g2
5 f8=Q    g1=Q
6 Qc8+    Ka5
7 Qc3+    Ka6
8 Qc4+    Kb6
9 a5+    Kxa5
10 Qb4+    Ka6
11 Qa4+    Kb6
12 Qa7+
again White wins 
the Queen.

78VAR-7
1 a3!    f5
2 gxf5    g4
3 f6        g3
4 f7        g2
5 f8=Q    g1=Q
6 Qc8+    Ka5
7 Qc3+    Kb6
8 a5+    Kb5
9 Qb4+    Kc6
10 Qb7+    Kd3
11 Qb6+
brutal but con-
vincing Queen 
exchange.

79VAR-1
1  f3    g3
2 f4        b6
3 Kf3
wins the luckless 
Pawn.

80VAR-1
1  c5        dxc5
2 bxc5    b6
3 d6        cxd6
3 c6
White wins with
earlier Queen than
Black.

80VAR-2
1 c5        Kg4
2 b5!        axb5
3 c6
decisive when Rook
Pawn is released.

81VAR-1
1 c5!    dxc5
2 d6    cxd6
3 Kxd6    b5
4 Kc6
and all Black's Pawns 
will fall.

81VAR-2
1 c5!    bxc5
2 Kb5    Kd7
3 Kxa5    c6
Black has a 
counter-play.

81VAR-3
1 c5!    bxc5
2 Kb5    Kd7
3 a4        c6+
4 dxc6    Kc7
5 b3
Black is helpless.

82VAR-1
1 a6      f1=Q
2 a7        Kg8
3 a8=Q+    Kf7
4 Qb7+    Kf8
5 Qe7+    Kg8
6 Qe8# 

82VAR-2
1 a6      f1=Q
2 a7        Kg8
3 a8=Q+    Kf7
4 Qb7+    Kxf6
5 Qg7+    Kf5
6 Qf7+
Queen falls in a 
gorgeous skewer.

83VAR-1
1 f5        gxf5
2 e6    fxe6
3 g6
does the trick to 
break through for
a Queen.

.83VAR-2
1 f5        e6
2 fxe6    fxe6
3 f4        Kb8
4 f5!        gxf5
5 g6
White wins after
first Queening.

84VAR-1
1 g3+    Kg5
2 g4        Kh4
3 Kg2    Kg5
4 Kh3    Kh6
5 Kh4    Kg6
6 g5    Kg7
7 Kh5    Kh7
8 g6+    Kg7
9 Kg5    Kg8
10 Kf6    Kf8
11 Ke6    Kg7
12 Kd7
and White wins.

84VAR-2
1 g3+    fxg3+
2 Kg2    Kh5
3 Kxg3    Kg5
4 f4+    exf4+
5 Kf3    Kg6
6 Kxf4    Kf6
7 e5+    dxe5+
8 Ke4    Kf7
9 Kxe5    Ke7
10 d6+    cxd6+
11 Kd5    Ke8
12 Kxd6    Kd8
13 c7+     Kc8
14 Kc6
allowing a draw 
by stalemate.

84VAR-3
1 g3+    fxg3+
2 Kg2    Kh5
3 Kxg3    Kg5
4 f4+    exf4+
5 Kf3    Kg6
6 Kxf4    Kf6
7 e5+    dxe5+
8 Ke4    Kf7
9 Kxe5    Ke7
10 d6+    cxd6+
11 Kd5    Ke8
12 Kxd6    Kd8
13 c7+     Kc8
14 Ke6    Kxc7
15 Ke7    Kb8
16 Kd6    Kb7
17 Kd7    Kb8
18 Kc6    Ka7
19 Kc7    Ka8
20 Kxb6    Kb8
21 Ka6    Kc7
22 Ka7
escorts the Pawn 
to the eighth
square.

85VAR-1
1 a4        Kg3!
2 a5        h5
3 a6        h4
4 a7        h3
5 a8=Q    h2#!
snookered by
one lousy tempo. 

85VAR-2
1 Kxg2    Kg5
2 a4    bxa3e.p.
3 bxa3    Kf6
4 a4        Ke7
5 a5        Kd8
6 a6        Kc8
7 a7        Kb7
Black wins the 
Pawn and the 
ending.

85VAR-3
1 f6!    gxf6
2 Kxg2     Kg5
3 a4     bxa3e.p.
4 bxa3     Kf5
5 a4         Ke5
6 a5    Kxd5
King catches the
Queen at a8
this time!

85VAR-4
1 f6!    gxf6
2 Kxg2     Kg5
3 a4     bxa3e.p.
4 bxa3     Kf5
5 a4         Ke5
6 c6     d6
7 a5        Kxd5
8 a6        Kxc6
King catches the
freshly Queened
Pawn again. 

85VAR-5
1 f6!    gxf6
2 Kxg2     Kg5
3 a4     bxa3e.p.
4 bxa3     Kf5
5 a4         Ke5
6 d6!    c6
7 a5        Kd5
8 a6
one step ahead 
White wins this
time.

86VAR-1
1 h6        Nd6
2 h7        Nf7+
3 Ke7    Ne5
4 h8=Q    Ng6+
winning the Queen. 

86VAR-2
1 h6        Nd6
2 h7        Nf7+
3 Ke7    Ne5
4 Kf6    Nd7+
5 Kg7 
Knight is kept at a
distance, and wins.

86VAR-3
1 h6        Nd6
2 h7        Nf7+
3 Ke7    Nh8
4 Kf8    Ke5
5 Kg7    Ke6
6 Kxh8    Kf7
and White has 
	painted himself
into a corner with
stalemate.

87VAR-1
1 e6!        Ne2+
2 Kf3    Nd4+
and forked winning 
of the Pawn.

87VAR-2
1 e6!        Ne2+
2 Kh3    Nf4+
winning the Pawn,
same way.

87VAR-3
	1 e6!    Ne2+
2 Kg4    Nc3
3 e7        Nd5
4 e8=Q    Nf6+
winning the Queen,
same way.

87VAR-4
1 e6!        Ne2+
2 Kh4    Nf4
3 e7        Ng6+
winning the Pawn-
have we learned
this lesson yet?.

87VAR-5
1 e6!        Ne2+
2 Kf2    Nc3
3 e7        Ne4+
4 Ke3    Nd6
stopping the Pawn, 
later capturing it 
with the King.

87VAR-6
	1 e6!    Ne2+
2 Kg2    Nf4+
winning the Pawn. 
How that Knight 
hops around!

88VAR-1
1 b7        Ne5+
2 Kd5    Nd7
3 Kd6    Nb8
4 Kc7    Na6+
5 Kb6    Nb8
6 Ka7    Nd7
and Black draws.

88VAR-2
1 Kd5     Ne3+
2 Kc5
White will Queen.

88VAR-3
1 Kd5    Nf6+
2 Kc6
White Queens again.

89VAR-0

90VAR-0

91VAR-1
1 a5        Bf8
2 Kd5    Bh6
3 g5+!    Kxg5
4 a6        Bf8
5 a7 
too late for Bishop
and Pawn will Queen.

92VAR-1
1 b7        Be5
2 Kd5    Bb8
the position draws.

92VAR-2
1 Kd5!    Be5
2 g3+    Kf5
3 g4+    Kf4
4 g5    Kf5
5 g6    Kf6
6 g7
and White wins.

92VAR-3
1 Kd5!    Be5
2 g3+    Kf5
3 g4+    Kf6
4 g5+    Kf5
5 g6        Bc3
6 b7        Be5
7 b8=Q    Bxb8
8 g7
White wins.

93VAR-1
1 Kf5    Bb6
2 Ke5    Ke2
3 h7    Kd3
Bishop reaches 
d4 pulling off a
stalemate.

93VAR-2
1 Kf5    Bb6
2 Ke4    Bd8
3 Ke5!    Be7
4 h7        Bxb4
5 Kd4    Ba3
6 Kc3
White controls   
diagonal where  
Pawn will Queen,
	and win.

94VAR-1
1 Kg7    Bb3
2 h5        Kd7
3 h6        Bc2
4 Kf6    Ke8
draws the position. 

95VAR-1
1 c5        Kg5
2 c6        Bd8
3 Ke5    Kh6
4 Ke6    Kg7
5 f6+    Kf8
Black has a draw.

96VAR-1
1 b5    Kd8
2 Kb7    Kd7
3 b6    Ba5!
4 Ka7    Bxb6+
5 Kxb6    Kc8
6 Ka7    Kc7 
7 a5    Kc8
8 a6    Kc7
and Black plasters
White against the 
Wall for a draw.

96VAR-2
1 b5        Kd8
2 Kb7    Kd7
3 b6        Ba5!
4 Ka7    Bxb6+
5 Kxb6    Kc8
6 a5        Kb8
and Black draws.

96VAR-3
1 a5!    Bxb4
2 a6
the Pawn cannot 
be stopped. 

97VAR-1
1 c7    Bc6+
2 Kb8    Bb7
3 b5+    Kxb5
4 c8=Q    Bxc8
5 Kxc8    Kxb6
and Black draws. 

97VAR-2
1 c7    Bc6+
2 Kb8    Bb7
3 c8=Q    Bxc8
4 Kc7    Bb7
5 b5+    Kxb5
6 Kxb7
Black has been 
swindled. 

97VAR-3
1 c7    Bc6+
2 Kb8    Bb7
3 c8=Q    Bxc8
4 Kc7    Kb5
5 Kxc8    Kxb6
the position draws.

98VAR-1
1 f5        Kg3
2 g5!    Bxh5
3 gxh6
wins at once

98VAR-2
1 f5        Kg3
2 g5!    Kg4
3 g6    Bd5
4 f6        Kxh5
5 f7
and White wins.

99VAR-0

100VAR-1
1 a7    Bg2
wrecking White's 
chances of a win.

100VAR-2
1 c6        dxc6
2 a7        c5+
3 Kxc5    Bg2
and Black draws.

100VAR-3
1 c6        dxc6
2 Kc5    Bc8
3 a7        Bb7
4 Kd6          Ba8
5 Ke7           c5
6 f7+
Pawn Queening
assures the win.

101VAR-1
1 c7        Rd6+
2 Kb7        Rd7 
followed by capturing 
the pinned Pawn draws. 

101VAR-2
1 c7        Rd6+
2 Kc5        Rd1 
3 Kc6         Rc1+
after capture of the skewered 
Pawn, draws. 

101VAR-3
1 c7        d6+
2 Kb5        Rd5+
3 Kb4        Rd4+
4 Kb3        Rd3+
5 Kc2        Rd4
6 c8=Q     Rc4+!
7 Qxc4
Black draws by stalemate.

101VAR-4
1 c7        Rd6+
2 Kb5        Rd5+
3 Kb4        Rd4+
4 Kb3        Rd3+
5 Kc2        Rd4
6 c8=R!   
White threatens to 
mate with 7 Ra8+.

102VAR-1
1 Kf7        Ra8
2g8=Q     Rxg8
3 Kxg8
and remaining Pawn becomes a Queen.

102VAR-2
1 Kf7         Rf2+
2 Ke6        Re2+
3 Kf5         Re8
4 Kf6        Kc7
5 Kf7        Kd7
6 g8=Q
wins for White.

102VAR-3
1 Kf7         Rf2+
2 Ke6       Re2+
3 Kf5         Re8
4 Ke4        Re2+
5 Kf3        Re8
6 Kf4        Rg8
7 Kf5        Rxg7
Rook gives up his life
for the second Pawn, 
draws for Black.

103VAR-1
1 Kcl        a3
2 Nc2+         Ka2
3 Nd4        Ka1
4 Kc2        a2
5 Nb3#

104VAR-1
1 Nc5        Kd2
2 Ne4+         Ke3
would threaten to capture the
Knight and then the Pawn.

105VAR-1
1 Kb4        Ka1
2 Kxa3
is a draw by stalemate.

106VAR-1
1 Kf5        Kg8
2 Kg6        Kh8
3 h7        a2
4 Nxa2
stalemate.

107VAR-1
1 Ng5+     Kg4
	draws easily.

107VAR-2
1 h3         Kg3
2 Ng5        Kf4
3 Ne4        Kf3
4 Kd4        Kg2
5 Ng5        Kg3
6 Ke3        Kg2
7 Kf4        Kf2
8 Kg4
	and White wins.

107VAR-3
1 h3        Kg3
2  Ng5        Kf4
3  Ne4        Kf3
4 Kd4        Kf4
5 Kd5        Kf5
6 Nf2        Kf4
7 Ke6        Kg3
8 Kf5        Kxf2
9 Kg4        Ke3
10 Kxh4     Kf4
11 Kh5        Kf5
12 Kh6        Kf6
13 Kh7        Kf7
14 h4        Kf8
15 h5        Kf7
16 h6        Kf8
17 Kg6 
allowing Black to 
reach g8 and a draw.

108VAR-1
1 Ng2!        Kh2
2 Nxh4     Kh3
3 Nf5        h5
4 Kgl        h4
5 gxh4
Rook Pawn marches
	up the board to Queen.

109VAR-1
1 d7        Bc6
2 Kd8         Bxd7
draws the position.

109VAR-2
1 Kd7!        Kd5
2 Kc7        Bc6
3 Ne4!        Kxe4
4 Kxc6
and the Pawn Queens.

109VAR-3
1 Kd7!        Kd5
2 Kc7        Bc6
3 Ne4!        Be8
4 Nf6+
winning the Bishop
and the game.

109VAR-4
1 Kd7!        Kd5
2 Kc7        Bc6
3 Ne4!        Bb5
4 Nc3+     Kc5
5 Nxb5          Kxb5
6 d7
White winning.

109VAR-5
1 Kd7!        Kd5
2 Kc7        Bc6
3 Ne4!        Ba4
4 Nc3+     Ke5
5 Nxa4
White wins Bishop
	and game.

110VAR-1
1 Nf5        Bc8
2 Ne7+    Kd7
3 Nxc8        
wins for White.

110VAR-2
1 Nf5        Ba8
2 Nd4+     Kc5
3 Ne6+         Kb5
4 Nc7+
White wins

110VAR-3
1 Nf5        Ba8
2 Nd4+    Kc5
3 Ne6+        Kc6
4 Nc7        Bb7
5 Nd5!        Kxd5
6 Kxb7
	wins for White.

110VAR-4
1 Nf5        Ba8
2 Nd4+    Kc5
3 Ne6+        Kc6
4 Nc7        Bb7
5 Nd5!        Bc8
6 Ne7+     Kd7
7 Nxc8         Kxc8
8 b7+
compels resignation and.
White wins

111VAR-1
1 Nc6        Bf1
2 b6        Ba6
3 Kd6        Ke3
4 Kc7        Ke4
5 Nb4
Bishop is evicted 
enabling the Pawn 
to advance next move.

111VAR-2
1 Nc6        Bf1
2 b6        Ba6
3 Kd6        Bb7
4 Kc7        Ba8
5 Nd8        Bf3
6 Nc6
Pawn Queens
in two more moves.

112VAR-1
1 Kb8          Kd8
2 Na5          Ba8!
3 Kxa8     Kc7
4 Nc4          Kc8
5 Nd6+       Kc7
6 Nb5+       Kc8
fails due to the fact 
that the Knight cannot
gain a move.

112VAR-2
1 Na5           Ba8
2 Kb8           Kd8
3 Nb7+     Kd7
4 Kxa8     Kc8
5 Nd6+       Kc
6 Ne8+        Kc8
and the stubborn 
King cannot be
ousted-- a draw.

113VAR-1  
1 Nb7+           Kc6
2 Nd8+     Kb6
3 Ne6 
followed by Queening 
with check wins. 

113VAR-2
1 Nb7+           Kc6
2 Nd8+     Kc5
3 Ne6+
	costs the Bishop. 
	
	113VAR-3
1 Nb7+           Kc6
2 Nd8+     Kb5
3 Ne6          Bg3
4 Kf6           Bh4+
5 Ng5 
White wins.

113VAR-4
1 Nb7+          Kc4
2 Nd6+     Kd5
3 Kf6!         Bg3
4 Kg5
White wins with a
beautiful variation.

113VAR-5
1 Nb7+          Kb6
2 d8=Q+
White wins.

113VAR-6
1 Nb7+     Kb4
2 Nd6         Bg5+
3 Ke8       Bf6
4 Nf5        Bg5
5 Ne7
blocking the Bishop's 
action on the diagonal.

113VAR-7
1 Nb7+     Kb4
2 Nd6           Be3
3 Nc8           Bf4
4 Nb6!           Bc7
5 Nd5+
winning the Bishop.

113VAR-8
1 Nb7+     Kb4
2 Nd6        Be3
3 Nc8        Bf4
4 Nb6!        Bg5+
5 Ke8        Kb5
6 Nc8       Bf6
7 Ne7
shuts the Bishop out.

113VAR-9
1 Nb7+     Kb4
2 Nd6        Be3
3 Nc8        Bf4
4 Nb6!        Kc5
5 Nd5!         Kxd5
6 d8=Q+
emphasize Pawn 
Queens with check.

113VAR-10
1 Nb7+     Kb4
2 Nd6        Be3
3 Nc8        Bf4
4 Nb6!        Kc5
5 Nd5!        Bg5+
6 Nf6
	is decisive.

114VAR-1
1 Ng7+!    Kf7
2 Nxf5        Kf6
3 h6 
wins for White. 

114VAR-2
1 Ng7+!    Nxg7
2 h6        Kf7
3 hxg7        Kxg7
drawing is fit
	retribution.

115VAR-1
1 Nd2        Kg7
2 Nc4        Nc2
3 b5        Ne1
4 b6        Nd3+
5 Kb5    Nf4
6 b7 
leaves Black helpless
so far from the action.

115VAR-2
1 Nd2        Kg7
2 Nc4        Nb1
3 b5        Nc3
4 b6        Na4+
5 Kc6        Nxb6
Knights capture is a
bloody draw.

115VAR-3
1 Nd2        Kg7
2 Nc4        Nb1
3 Kd4!        Kf7
4 b5        Ke7
5 b6        Kd7
6 Kc5        Nc3
7 Ne5+        Ke6
8 b7        Na4+
9 Kb4 
White wins.

116VAR-1
1 Kb8        Kb5
2 a7        Nc6+
wins the Pawn 

116VAR-2
1 Kb8        Kb5
2 Nb4         Nc6+
3 Kc7        Nxb4
4 a7         Nd5+
5 Kc8          Nb6
draws
	
	116VAR-3
1 Kb8        Kb5
2 Nb4         Nc6+
3 Kb7         Na5+
4 Ka7        Kxb4
5 Kb6        Nc4+
6 Kc7        Ka5
7 a7        Nb6
	and Black draws.

116VAR-4
1 Ka7!        Nc6+
2 Kb6        Ne7
3 Kb7        
White wins with un-
challenged Pawn.

116VAR-5    
1 Ka7!        Nc6+
2 Kb6        Kd5
3 Nb4+    Nxb4
4 a7
White wins again.  

116VAR-6  
1 Ka7!        Kb5
2 Nb4!        Kc5
3 Kb8        Kxb4
4 a7        Nc6+
5 Ka8        Nxa7 
Pawn and Knight
removed for a draw!

116VAR-7  
1 Ka7!        Kb5
2 Nb4!        Kc5
3 Kb8        Kxb4
4 Kc7         Ne6+
5 Kb6
and Pawn cannot 
be stopped.

116VAR-8  
1 Ka7!        Kb5
2 Nb4!        Ka5
3 Kb8        Nc6+
4 Kb7        Nd8+
5 Kc7        Ne6+
6 Kb8        Kb6
7 a7        Nc7
8 Nd5+
White wins.

117VAR-1  
1 Ne6        Kd5
2 Nf8        Ne5
3 b8=Q    Nc6+
Queen comes off the board.

117VAR-2 
1 Ne6        Kd5
2 Nf8        Ne5
	3 Ka8        Nc6
4 Nd7        Ke6
5 Nb6        Kf5
6 Nc8        Ke6
7 Na7
springing the Pawn free.

117VAR-3 
1 Ne6        Kd5
2 Nf8        Ne5
3 Ka8        Nc6
4 Nd7        Kd6
5 Nb6        Kc7
6 Nd5+    Kd6
7 Nb4
	wins for White.

118VAR-1
1 Nf6        Nd7+
2 Ka8        Nb6+
3 Kb8        Nd7+
4 Kc8        Nb6+
	etc., Black has a clever defense 
against perpetual check with:

118VAR-2
1 Nf6        Kc5
2 Kb7        Kb5
3 Nd5 
evicts the Knight from holding 
back the Pawn and wins.

118VAR-3
1 Nf6        Na8!
2 Kxa8        Kc7
3 Nd5+    Kc8
4 Ne7+    Kc7
and Black draws.

118VAR-4
1 Nf6        Na8!
2 Nd5!        Kd7
3 Kxa8    Kc8
another Black draw.

118VAR-5
1 Nf6        Na8!
2 Nd5!        Kd7
3 Kb7        Kd6
4 Kxa8    Kxd5
5 Kb8
White winning.

119VAR-1
1 Kf5        Ka7!
2 Nd7        Nxd7
3 e6        Nb6
4 e7        Nc8
5 e8=Q    Nd6+
Black wins the Queen. 

119VAR-2
1 Kf5        Ka7
2 Kf6        Kb6   
3 Kg7        Kxc5
4 Kxf8        Kd5
Black captures the 
Pawn and draws. 

119VAR-3
1 Kf5        Ka7
2 Kg5        Kb6
3 Nd7+    Nxd7
4 e6        Nc5
5 e7        Ne6+
6 Kf6        Nc7
Pawn can never 
move on to Queen.

119VAR-4
1 Kg5!        Ka7 
2 Kh6         Kb6
3 Kg7        Kxc5 draws. 

119VAR-5
1 Kg5!        Ka7 
2 Kh6         Kb6
3 Nd7+    Nxd7
4 e6        Nf6
5 e7        Ng8+
and Black draws.

119VAR-6
1 Kg5!        Ka7 
2 Kh6         Kb6
3 Nd7+    Nxd7
4 e6        Nf6
5 Kg6         Nd5
and Black draws again.

119VAR-7
1 Kg5!        Ka7
2 Kf5!        Kb8
3 Kf6        Nh7+
4 Kg6        Nf8+
5 Kg7
Knight is magnificently
	trapped. 

119VAR-8
	Kg5!        Ka7
2 Kf5!        Kb8
3 Kf6        Kc8
4 Kg7        Nd7
5 Nxd7    Kxd7
6 Kf7
the Pawn will Queen.

119VAR-9
1 Kg5!        Ka7
2 Kf5!        Kb6
3 Nd7+!    Nxd7
4 e6        Kb7
5 e7        Kc7
6 e8=Q
White wins.

119VAR-10
1 Kg5!        Ka7
2 Kf5!        Kb6
3 Nd7+!    Nxd7
4 e6        Nc5
5 e7        Nb7
6 e8=Q    Nd6+
removes the Queen.

120VAR-0





121VAR-1
1 f5        Ne3
2 f6        Nf5
3 f7        Ne7+
4 Kb7        Ng6
5 Nxe5        Nxe5
6 f8=Q
and White wins.

121VAR-2
1 f5        Ne3
2 f6        Nf5
3 f7        Ne7+
4 Kb7        Ng6
5 Nxe5        a3
6 Nxg6    a7
7 f8=Q        a1=Q
8 Qa8+
Black's Queen is
removed.

121VAR-3
1 f5        Ne3
2 f6        Nf5
3 f7        Ne7+
4 Kb7        Ng6
5 Nxe5        Nf8
6 Nc6#!    

122VAR-1  
1 bxa7        Rb3+
2 Kc7        Rc3+
3 Kd6        Rc8
solves Black's problems. 

122VAR-2
1 bxa7        Rb3+
2 Kc7        Rc3+
3 Kd8        Rxh3
White does not dare 
Queen his Pawn.

123VAR-1
1 Nxh6    gxh6
2 Kf6        Kg8
3 g7        Kh7
4 Kf7
stalemates Black.  

123VAR-2
1 Nxh6    gxh6
2 Kf6        Kg8
3 g7        Kh7
4 g8=Q+    Kxg8
5 Kg6
and White can win
the remaining Pawn. 

124VAR-1
1 Nb4+!    axb4
2 e7        d2
3 e8=Q        d1=Q
4 Qc6+        Ka5
5 Kb7!        b3
6 a3           Qd3
7 Qc5+        Ka4
8 Qb4#

125VAR-1
1 b7        Nd6+
2 Kd5        Nxb7
3 e7        Kf7
Black has a draw.

125VAR-2
1 b7        Nd6+
2 Kd4!        Nxb7
3 Kd5        Nc5
4 e7        Na6
5 e8=Q        Nc7+
winning the Queen. 

125VAR-3
1 b7                 Nd6+
2 Kd4!        Nxb7
3 Kd5        Nc5
4 e7        Na6
5 Kd6         Kf7
6 Nd8+    Ke8
7 Ne6        Kf7
8 Ng7        Nc7
9 Kd7        Kf6
10 Ne8+    Nxe8
11 Kxe8
	and White wins.

126VAR-1
1 Ke7        Kg5
2 Kxf8        Kxh5
3 Kg7        Kg4
4 Ne4        Kf3
5 Kg6        Kg4
6 Kf6        Kf3
7 Kf5
and White wins.

126VAR-2
1 Ke7        Nf7
2 Kf7        Nf6
3 h6        Ke5
4 Nd7+
White wins. 

126VAR-3
1 Ke7        Nh7
	2 Kf7        Nf6
3 h6        Kg5
4 Ne4+        Nxe4
5 h7
wins for White.

127VAR-1
1 g6          Kf6
2 g7          Kf7
3 Nd5          Kg8
4 Nf6          Kf7
5 g8=Q+ 

127VAR-2
1 g6        Kf6
2 g7        Kf7
3 Nd5        Nd7
4 Ne7        Nf6
5 g8=Q+    Nxg8
6 h7
and White Queens
and wins.

127VAR-3
1 g6        Kf6
2 g7        Kf7
3 Nd5        h7
4 Ne7        Nf6
5 g8=Q+    Nxg8
6 h7
	and White wins.

127VAR-4
1 g6        Kf6
2 g7        Kf7
3 Nd5        Ng6
4 Nf6        Ne7
5 g8=Q+    Nxg8
6 h7
White wins again.

128VAR-1
1 Nd4+     Kc5
2 Kh1!        Bf8
3 Ne6+
wins the Bishop.

128VAR-2
1 Nd4+     Kc5
2 Kh1!        Bg7
3 Ne6+ 
wins the Bishop.

128VAR-3
1 Nd4+     Kc5
2 Kh1!        Bg5
3 Ne6+
wins the Bishop.

128VAR-4
1 Nd4+     Kc5
2 Kh1!        Bf4
3 Ne6+
wins the Bishop.

128VAR-5
1 Nd4+     Kc5
2 Kh1!        Bd2
3 Nb3+
wins the Bishop.

128VAR-6
1 Nd4+     Kc5
2 Kh1!        Bc1
3 Nb3+
wins the Bishop.

128VAR-7
1 Nd4+     Kc5
2 Kh1!        Kd6
3 Nf5+
wins the Bishop.

128VAR-8
1 Nd4+     Kc5
2 Kh1!        Kd5
3 a6
passed Pawn becomes 
a Queen.

129VAR-1
1 c6        dxc6
2 a6              c5
3 Ne5
prevents Bishop from 
getting on the diagonal
with 4 Nc6!.

129VAR-2
1 c6        dxc6
2 a6              Bf3
3 Ng5              Bd5
4 Ne6                  Kd7
5 Nc5+    Kc7
6 a7
	leaves Bishop helpless
and Pawn Queens.

129VAR-3
1 c6        dxc6
2 a6              Bf3
3 Ng5              Bd5
4 Ne6                  c5
5 Nc7+        Kd7
6 Nxd5        Kc6
7 Kg2        c4
8 Kf2        c3
9 Ke3
White overtakes the Pawn.

129VAR-4
1 c6        dxc6
2 a6              Bf3
3 Ng5              Bd5
4 Ne6                  c5
5 Nc7+        Kd7
6 Nxd5        Kc8
7 Nb6+        Kc7
8 a7
followed by Queening.

130VAR-1
1 c5        Bd5
2 c6        bxc6
3 b7
lets the Pawn escape.

130VAR-2
1 c5        Bb1
2 Ne6!        Be4
3 Ng5+    Kf4
4 Nxe4    Kxe4
5 c6
etc., with White 
Queening first, and wins.

131VAR-1
1 f7        Rxa6+
2 Nf6        Ra8
3 Ne8        Ra6+
4 Kg7        Ra7
next, sacrificing the Rook
for the Pawn, draws, i.e.
K+N v. K scenario.

131VAR-2
1 f7        Rxa6+
2 Nf6        Ra8
3 Ne8        Ra6+
4 Kf5        Ra1
5 Kd6        Rf1
with or without check,
a Rook for a Pawn,
draws for Black.

132VAR-1
1 b7        Rf8
2 Nb4+        Kc4
3 Nc6        Kc5
4 b8=Q    Rxb8
5 Nxb8    Kd5
6 g4        Ke5
7 g5        Kf5
8 Nd7        Kf4
9 Nf6 
and White wins.

132VAR-2
1 b7        Rf8
2 Nb4+        Ke4
3 Nc6        Kf4
4 b8=Q    Rxb8
5 Nxb8    Kg3
Black captures the 
Pawn next move. 
and draws.

132VAR-3
1 b7        Rf8
2 Nb4+        Ke4
3 Nc6        Kf4
4 g4!        Kxg4
White stopped cold.

132VAR-4
1 b7        Rf8
2 Nb4+        Ke4
3 Nc6        Kf4
4 g4!        Rf6+
5 Kxh7    Rf7+
and White loses the 
ambitious Queen 
Knight Pawn.

132VAR-5 
1 b7        Rf8
2 Nb4+        Ke4
3 Nc6        Kf4
4 g4!        Rf6+
5 Kg7         Rg6+
6 Kxh7
	and White wins.

132VAR-6
1 b7        Rf8
2 Nb4+        Ke4
3 Nc6        Kf4
4 g4!        Kxg4
5 Kg7         Re8
6 Kf7        Rh8 
7 Ke7!        Rg8
8 Nd8        Rg7+
	9 Kf6        Rxb7
Rook for Pawn swap
draws easily.

133VAR-1
1 Kg4        Kc8
2 Kh5        Kd8
3 Ng7!        Bxg7
4 h8=Q+!
unexpected, and the 
only way to win. 

133VAR-2
1 Kg4        Kc8
2 Kh5        Kd8
3 Ng7!        Bxg7
4 Kg6        Bh8
5 Kf7        Kd7
6 Kg8        Ke7
7 Kxh8    Kf7
and White is 
stalemated.

134VAR-0

135VAR-1
1 h6        Kf7
2 Bh7!        Kg8
3 Kf4         Kh8
which would assure 
Black of a draw.

136VAR-1
1 Be6!        Ke7
2 h6        Kf6
3 Bf7        Kf7
where the Pawn
is attacked.

136VAR-2
1 Be6!        Ke7
2 h6        Kf6
3 Bf5!        Kf7
4 Kf4           Kg8
and a sure draw.

137VAR-1
1 Kf7        h5
2 Bf6        Kb3
3 Kg6        h4
4 Bxh4        Kxb2
and Black has a King-
Bishop v. King draw.

137VAR-2
1 Kf7        h5
2 Bf6        Kb3
3 Kf5
would close in 
on the Pawn.

137VAR-3
1 Kf7        h5
2 Bf6        h4
3 Kd5        Kb3
4 Ke4
Rook Pawn is doomed.

138VAR-1
1 Bc3        Kc8
2 Ka4         Kb8
wth Black a 
certain draw.

138VAR-2  
1 Bc3        Kc8
2 Be5        b2
3 Bxb2    Kb8
White cannot possibly 
force a win. 

138VAR-3 
1 a7        b2
2 a8=Q        b1=Q
3 Qb7+    Ke8
4 Qe7# 

138VAR-4 
1 a7        b2
2 a8=Q        b1=Q
3 Qb7+    Kd8
4 Be7+
discovering an attack 
on the Queen.

138VAR-5 
1 a7        b2
2 a8=Q        b1=Q
3 Qb7+    Ke6
4 Qe7+        Kf5
5 Qh7+
winning the Queen.

138VAR-5 
1 a7        b2
2 a8=Q        b1=Q
3 Qb7+    Ke6
4 Qe7+        Kd5
5 Qd6+    Ke4
6 Qg6+
wins the Queen.

139VAR-1  
1 Bc5        Kf7
2 a4        Ke8
3 a5        Kd8
4 Bd6        Kc8
5 a6        Kd8
6 a7
and White wins. 

139VAR-2
1 Bc5        Kf7
2 a4        Ke6
3 a5        Kd5
4 a6        Kc6
5 Kg2        Kc7
aiming for b8 and 
a sure draw. 

139VAR-3
1 Bc5        Kf7
2 a4        Ke6
3 a5        Kd5
4 a6        Kc6
5 Kg2        Kc7
6 Ba7        Kc6
7 Bb8        Kb6
8 a7        Kb7
White must concede 
the draw.

140VAR-1
1 Bc5        Ka5
2 Kb7        Kb5
3 Bb6!        Kc4
4 Kc6        Kd3
5 Kb5        Ke4
6 Kxa4    Kd5
7 Kb5        Kd6
8 Ka6        Kc6
9 a4        Kd7
10 Kb7
Black is kept out  
with plenty of protection 
for a Pawn Queening.

141VAR-1  
1 c8=Q    g1=Q+!
2 Qc5+
Black losing the Queen

141VAR-2
1 Bd4!        Kxd4
2 c8=Q    g1=Q
3 Qc5+
again Black forfeits 
His Queen.

141VAR-3
1 Bd4!        g1=Q
2 Bxg1        a2
3 c8=Q    a1=Q
would have drawn 
for Black.

142VAR-1
1 Bb1        h4
2 Bf5        d6
3 Bh3        Kd3
4 Bf1+        Kd4
5 Kb5        Kc3
6 Kc6        h3
7 Kd5        h2
8 Bg2        Kd3
captures Queen Pawn 
only at the expense of 
his own Pawn.

142VAR-2
1 Bb1        h4
2 Bf5        d6
3 Bh3        Kd3
4 Bf1+        Kd4
5 Kb5        Kc3
6 Kc6        h3
7 Kxd6    h2
8 Bg2        Kxc4
and a draw.

143VAR-1
1 Bg1+         c3
2 bxc3        Kc4
invaluable Pawn comes 
off the board, enabling 
Black to draw.

143VAR-2
1 Bg1+        Kb4
2 Bd4        Kb3
3 Bc3        b4
4 Bd4        c3
Black forces the draw.

144VAR-1
1 Kc6        Ka6
2 Be3        Ka5
3 Bc5        Ka6
4 Bb6        b4
5 axb4        a3
6 b5#

144VAR-2
1 Kc6        Ka8
2 Kxb5??
after which the King-
Bishop v. King scenario 
draws.

145VAR-1
1 Bf3+        Kg1
2 Bh1!        Kxh1
3 Kf1        d5
4 exd5        e4
5 d6        e3
6 d7        e2+
7 Kxe2        Kg1
8 d8=Q    h1=Q
9 Qg5+?    Qg2+
forcing a draw.

145VAR-2
1 Bf3+        Kg1
2 Bh1!        Kxh1
3 Kf1        d5
4 exd5        e4
5 d6        e3
6 d7        e2+
7 Kxe2        Kg1
8 d8=Q    h1=Q
9 Qd4+    Kh2
10 Qh4+    Kg1
11 Qf2#

146VAR-1
1 Bb1!        Kb3
2 Kc6        Kb2
3 d7        Kxb1
4 d8=Q    a2
5 Qb6+    Ka1
6 Qxa5    Kb2
7 Qb4+    Kc2
8 Qa3        Kb1
9 Qb3+    Ka1
10 Qc2        f4
11 Qc1#

146VAR-2
1 Bbl!        f4
2 Kc 6        f3
3 Kc5!        Kb3
4 d7        f2
5 d8=Q        f1=Q
6 Qd5+        Kb2
7 Qa2+        Kc3
8 Qc2# 

146VAR-3
1 Bbl!        f4
2 Kc 6        f3
3 Kc5!        Kb3
4 d7        f2
5 d8=Q        f1=Q
6 Qd5+        Kc3
7 Qd4+        Kb3
8 Qa4+!        Kb2
9 Qc2+        Ka1
10 Qa2#

147VAR-1
1 Bf7                Kb5
2 Be8+               Ka5
3 Bh5        Ne6+
4 Kc6        Nd8+
5 Kc5        Ne6+
6 Kc6        Nd8+
Black has a perpetual
check- aka draw.

147VAR-2
1 Bf7        Kb5
2 Be8+            Ka5
3 Bd7        Ka6
4 Bh3           Nb7
5 Bg2
White wins.

147VAR-3
1 Bf7        Kb5
2 Be8+            Ka5
3 Bd7        Ka6
4 Bh3           Nb7
5 Bf1+
White wins.

147VAR-4
1 Bf7        Kb5
2 Be8+            Ka5
3 Bd7        Ka6
4 Bh3           Ka5
5 Bg4        Kb5
6 Be2+     Kc5
7 Bc4
leaves Black helpless.

148VAR-1
1 Kg5        Nf2
2 h4!        Ne4+
3 Kg6        Nf2
4 h5        Ng4
5 Kg5
Knight must retreat,
King can no longer 
move and protect it.

151VAR-1
1 Bg5        Bb2
2 Bf6
wins instantly.

151VAR-2
1 Bg5        Bf8
2 Kf 6!        Kg4
3 Bd2        Kh5
4 Kf7
Black is out 
of moves. 

151VAR-3
1 Bg5        Bf8
2 Kf 6!        Kg4
3 Bd2        Bc5
4 Bc3        Bf8
5 Kf7        Kf5
6 Bd2
and White wins.

152VAR-1
1 Kh7!        Bb2
2 Bf4        Bd4
3 Bh6+        Ke8
4 Bg7        Be3
5 Bb2         Bg5?
6 g7
with White winning.

153VAR-1
1 Bg7        Bg5
2 Bh6        Bf6
3 Bg5        Bg7
4 Be7!        Bc3
5 h6        Bd4
6 Bf6
White wins.

153VAR-2
1 Bg7        Bg5
2 Bh6        Be7
3 Be3        Bf8
4 Bd4        Kf4
5 Bg7
is conclusive.

153VAR-3
1 Bg7        Bg5
2 Bh6        Be7
3 Be3        Bf8
4 Bd4        Kh4
5 Be5        Kg4
6 Bf6        Kf4
7 Bg7
Pawn now has 
a clear road ahead.

154VAR-1
1 Bh4        Be5
2 Bf6        Bh2
3 Bd4        Bg3
4 Ba7        Bf4
5 Bb8        Be3
6 Bh2        Ba7
7 Bg1
and White wins.
    
154VAR-2
1 Bh4        Kb6
2 Bf2+        Kc6
3 Ba7        Bg3
4 Bb8        Bf2
5 Bh2        Ba7
6 Bg1
again Bishop is
cornered, White wins.

154VAR-3
1 Bh4        Kb6
2 Bf2+        Ka6
3 Bc5!        Bg3
4 Be7        Kb6
 5 Bd8        Kc6
6 Bc7
is decisive in draw.

154VAR-4
1 Bh4        Kb6
2 Bf2+        Ka6
3 Bc5!        Bf4
4 Be7        Kb6
5 Bd8+        Kc6
6 Bg5
a draw.

154VAR-5
1 Bh4        Kb6
2 Bf2+        Ka6
3 Bc5!        Be5
4 Be7        Kb6
5 Bd8+        Kc6
6 Bf6
another draw.

155VAR-1  
1 Ke5        c6
2 Bxc6        Kc7
3 Bg2        Na8!
4 Bxa8        Kb6
5 Bg2        Kxa7
 and Black has a King
v. King-Bishop draw.

155VAR-2  
1 Bc6!        Kd8
2 Kf7        Kc8
3 Ke8
easy win with the
Knight out of there.

155VAR-3  
1 Bc6!        Kd8
2 Kf7         Nc8
3 a8=Q
and stalemate!

155VAR-4
1 Bc6!        Kd8
2 Ke5        Ke7
giving Black the 
opposition and a draw. 

156VAR-1  
1 g5+        Ke7
2 g6        Ke8
3 g7        Bf8
4 Bf5         Bxg7
5 Kxg7
sacrificing the Bishop 
for the precious pawn;
resulting in a draw.

157VAR-1
1 h7        e4
ruins White’s 
winning chances.

157VAR-2
1 Ba7!        Bxa7
2 h7
wins for White.

157VAR-3
1 Ba7!        Ba1
2 Kb1        Bc3
3 Kc2        Ba1
4 Bd4!    exd4
5 Kd3 
blocks Bishop's path 
assuring Rook Pawn 
of Queening.

157VAR-4
1 Ba7!        Ba1
2 Kb1        Bc3
3 Kc2        Ba1
4 Bd4!    Bxd4
5 Kd3        Kg5
6 h7        e4+
7 Kxd4
Pawn cannot be 
over-taken.

157VAR-5
1 Ba7!        Ba1
2 Kb1        Bc3
3 Kc2        Ba1
4 Bd4!        Bxd4
5 Kd3        Ba1
6 h7        e4+
Bishop will prevent 
Pawn from Queening.

158VAR-1
1 Bf2+    Kh5
2 g4+        Kh6
3 Kf6        Bh7
4 Be3#

158VAR-2
1 Bf2+    Kh5
2 g4+        Kh6
3 Kf6        Kh7
4 g5        Kh8
5 Bd4        Bh7
6 Kxf7#

158VAR-3
1 Bf2+    Kh5
2 g4+        Kh6
3 Kf6        Kh7
4 g5        Kh8
5 Bd4        Kh7
6 Ba1        Kh8
7 g6!        Bh7
8 Kxf7#

159VAR-1
1 e5        Bb7
2 e6        Bc8
3 e7        Bd7
4 Kf8        c5
wins for Black!

159VAR-2
1 e5        Bb7
2 e6        Bc8
3 e7        Bd7
4 Bh3!    Be8
5 Kf8        Bh5
6 Be6     f3
giving Black a draw.

160VAR-1
1 c6        Bd5
2 Bc4         g3
3 Bxd5+
and the Knight Pawn 
is harmless. 

160VAR-2
1 c6        Bd5
2 Bc4        Bxc4
3 c7
and White wins.

160VAR-3
1 c6        Ba4
2 c7        Bd7
3 Ke7        Bc8
4 Kd8        Bb7
5 Bg2    Bxg2
6 c8=Q 
White wins.

161VAR-1
1 d7        Rg2+
2 Kf7        Rf2+
3 Ke7        Re2+
4 Kd6
a the Rook is out 
of checks.

162VAR-1
1 b7        Rf5+
2 Be5        Rf8
3 Bd6+ 
a beaut of a fork 
wins the Rook.

162VAR-2
1 b7        Rd3+
2 Ke6!    Rd8
3 Bc7        Rh8
4 Be5        Re8+
5 Kf7        Rd8
6 Bc7        Rh8
7 Bd6+
leads by a transposition 
of moves into the actual play.

163VAR-1
1 c7        Rg5+
2 Kd4    Rg4+
3 Kc3        Rg3+
4 Kb2        Rg2+
5 Ka3
and White wins.

163VAR-2
1 c7        Re2+
2 Kf6        Re8
3 Ba4!        Rg8
4 Kf7        Rh8
5 Kg7        Rc8
6 Bd7+
Wins the Rook.

164VAR-1
1 a7        Rh8
2 Bf6+
wins the Rook.

164VAR-2
1 a7        Rf5+
2 Ke2     Rf8
3 Bf6+    Kc5
4 Bf7+
again the Rook goes.

165VAR-1
1 b7        Ra5+
2 Kxe6    Rb5
3 Bc6+    Kd8
4 Bxb5    Kc7
5 Ba6
and White wins. 

165VAR-2
1 b7        Ra5+
2 Kxe6    Ra6+
3 Kd5    Rb6
4 Kd4    Rxb7
Black draws with
King v. King-Bishop.
scenario.

165VAR-3
1 b7        Ra5+
2 Kd6!    Ra6+
3 Bc6+
Pawn will advance.

165VAR-4
1 b7        Ra5+
2 Kd6!    Rb5!
3 Bc6+    Kd8
4 Bxb5    Bc8!
5 b8=Q
stalemates Black.
There’s a better
move. 

165VAR-5
1 b7        Ra5+
2 Kd6!    Rb5!
3 Bc6+    Kd8
4 Bxb5    Bc8!
5 b8=B!    Bb7
6 Bc7+    Kc8
7 Bd7#

166VAR-1
1 g7        Be6+
2 Kxe6     Kxg7
Black draws against 
the Rook Pawn and 
a Bishop that does 
not control the Pawn's 
Queening square. 

166VAR-2
1 Kf6        Be8
2 h4        Bxg6
3 Bxg6
Black draws by stalemate. 

166VAR-3
1 Bg8        Be6+
2 Kxe6    Kxg6
Black gets to the corner
and draws.

166VAR-4
1 Kg8!        Bf5
2 g7!        Bxh7+
3 Kh8        Kg6 
4 h4        Kh6
5 h5
Black is in zugzwang,
though for every 
he move loses.

166VAR-5
1 Kg8!        Be6+
2 Kh8        Bd5
3 g7        Be6
4 Bg8        Bf5
5 Bd5        Bh7
6 Be4        Bxe4
7 g8=Q
White wins.

166VAR-6
1 Kg8!        Be6+
2 Kh8        Bf5
3 g7!        Bxh7
4 h4        Kg6
5 h5+        Kh6
White who is in zugzwang, 
must give up both Pawns
and draws.

167VAR-1
1 g6         Bc2
2 f5        Bb3
3 Kg5        Bc2
4 f6        Bb3
5 Bb4+    Ke8
6 Kh6          Bc2
7 Kg7          Bd1
8 f7+            Kd8
9 f8=Q
Black being unable 
meanwhile to do 
anything but watch

167VAR-2
1 g6         Bc2
2 f5        Bb3
3 Kg5        Bc2
4 f6        Bb3
5 Bb4+    Ke8
6 Kf4        Bc2
7 f7+
White  wins at once.

167VAR-3
 1 g6         Bc2
2 f5        Bb3
3 Kg5        Bc2
4 f6        Bb3
5 Bb4+    Ke8
6 Kf4        Bc4
7 Ke5        Bb3
8 Kd6        Bc2
9 f7+        Kg7
10 Ke7
 is decisive for White  
win at once again.

167VAR-4
1 g6         Bc2
2 f5        Bb3
3 Kg5        Bc2
4 f6        Bb3
5 Bb4+    Kg8
6 Kf4        Bc4
7 Ke5        Bb3
8 Kd6        Bc2
9 f7+        Kg7
10 Ke7
is decisive.

168VAR-1
1 Bc4+    Ke7
2 f5        Bg7
3 Kf4        Bh8
Black has an easy 
draw. 

168VAR-2
1 Bc4+    Ke7
2 f5        Bg7
3 Kf4        Bh8
4 e6
White has no winning 
chance. 

168VAR-3
1 Bc4+    Ke7
2 f5        Bg7
3 Kf4        Bh8
4 f6+
 allows the Bishop 
 to sacrifice itself for
 the two Pawns.

168VAR-4
1 Bc4+    Ke7
2 Ke4!    Bg7
3 Kf5        Bh6
4 Kg4!    Bg7
5 Kg5
 leads to the actual 
 play.

168VAR-5
1 Bc4+    Ke7
2 Ke4!    Bg7
3 Kf5        Bh6
4 Kg4!    Ke8
5 f5        Ke7
6 f6+        Ke8
7 e6
White wins easily.

168VAR-6
1 Bc4+    Ke7
2 Ke4!    Bg7
3 Kf5        Bh6
4 Kg4!    Bf8
5 Kg5        Bg7
6 Kg6        Kf8
7 Kh7
wins the Bishop.

168VAR-7
1 Bc4+    Ke7
2 Ke4!    Bg7
3 Kf5        Bh6
4 Kg4!    Bf8
5 Kg5    Bg7
6 Kg6        Bh8
7 Kh7
 cornering the Bishop 
literally and figuratively. 

169VAR-1
1 Ke2        Kg4
2 f3+        Kg3
time will be lost 
evicting the King.

169VAR-2
1 Ke2        Kg4
2 Be1        Bd6
3 f3+        Kf4
4 g3+        Kf5
5 g4+        Kf4
6 Bd2+    Kg3
7 g5        Be5
8 g6        Bg7
9 Bh6
10 g7
White wins. 

169VAR-3
1 Ke2        Kg4
2 Be1        Bd6
3 f3+        Kf4
4 g3+        Kf5
5 g4+        Kf4
6 Bd2+    Kg3
7 g5        Be5
8 g6        Bg7
9 Ke3        Bh6+
10 Ke4        Bg7
11 Bf4+    Kg2
12 Be5
is decisive. 

169VAR-4
1 Ke2        Kg4
2 Be1        Bd6
3 f3+        Kf4
4 g3+        Kf5
5 g4+        Kf4
6 Bd2+    Kg3
7 g5        Be5
8 g6        Bg7
9 Ke3        Kh4
10 Ke4    Kh5
11 Kf5    Bd4
12 Bg5    Bg7
13 Bf6    Bf8
14 g7 
wins for White.

169VAR-5
1 Ke2        Kg4
2 Be1        Bd6
3 f3+        Kf4
4 g3+        Kf5
5 g4+        Ke6
6 Kd3        Kd5
7 Bd2        Bc7
8 f4        Bb6
9 Bc3        Bc5
10 g5        Bb6
11 g6        Ke6
12 Ke4    Bd8
13 f5+        Ke7
14 Kd5    Bc7
15 f6+    Kf8
16 Bb4+    Kg8
17 f7+
and White wins.

170VAR-1
1 Bf1     Bg4
2 h4         Bf5
3 Kf2     Bg4
4 Ke3     Be6
5 Kf4     Bd7
6 g4         Be8
7 Be2     Bf7
8 g5+     Kh7
which results in
a Rook Pawn- Bishop
of the wrong color controlling
the Pawn’s Queening square.

170VAR-2
1 Bf1        Bg4
2 h4        Bf5
3 Kf2    Bg4
4 Ke3    Be6
5 Kf4    Bd7
6 Bd3    Bh3
7 Bf5        Bf1
8 g4        Be2
9 h4        Bxg4
forcing a draw.

170VAR-3
1 Bf1        Bg4
2 h4        Bf5
3 Kf2    Bg4
4 Ke3    Be6
5 Kf4    Bd7
6 Bd3    Bh3
7 Bf5        Bf1
8 g4         Be2
9 g5+     Kg7
10 Bg4     Bxg4
11 h5
and a draw.

170VAR-4
1 Bf1        Bg4
2 h4        Bf5
3 Kf2    Bg4
4 Ke3    Be6
5 Kf4    Bd7
6 Bd3    Bh3
7 Bf5        Bf1
8 g4         Be2
9 g5+     Kg7
10 g6        Kh6
11 Ke5    Bh5
12 Kf6    Bxg6
13 Bxg6
stalemate.

170VAR-5
1 Bf1        Bg4
2 h4        Bf5
3 Kf2        Bg4
4 Ke3        Be6
5 Kf4        Bd7
6 Bd3        Bh3
7 Bf5        Bf1
8 g4        Be2
9 g5+        Kh5
10 Kg3!        Bd1
11 Be4        Bb3
12 Bf3+        Kg6
13 Kf4        Bf7
14 h5+        Kg7
15 Ke5        Bb3
16 Be4        Bf7
17 h6+        Kh8
18 Kf6        Bh5
19 g6        Bxg6
and a draw. 

170VAR-6
1 Bf1        Bg4
2 h4        Bf5
3 Kf2        Bg4
4 Ke3        Be6
5 Kf4        Bd7
6 Bd3        Bh3
7 Bf5        Bf1
8 g4        Be2
9 g5+        Kh5
10 Kg3!    Bd1
11 Be4        Bb3
12 Bf3+    Kg6
13 Kf4        Bf7
14 h5+        Kg7
15 Ke5        Bb3
16 Be4        Bf7
17 h6+        Kh8
18 Kf6        Bh5
19  Bg6
Driving off Black's
Bishop, since it 
blocks the Knight Pawn.

171VAR-1
1 f4+        Kd6
2 f5        Ke5
3 d4+        Kf6
4 Kf4        Bb3
5 Bc6        Bc2
6 Bd7        Bb3
7 Ke4        Bc2+
8 Kd5        Bxf5
9 Bxf5    Kxf5
10 Kc6
wins for White.

171VAR-2
1 f4+        Kd6
2 f5        Ke5
3 d4+        Kf6
4 Kf4        Bb3
5 Bc6        Bc2
6 Bd7        Bb3
7 Ke4        Bc4
8 d5        Bb3
9 Be6        Bc4
10 Kd4        Ba2
11 Kc5    Bb3
12 d6
White wins.

172VAR-1
1 Kg5        Kxe7
2 Kxf5    Kf7
3 e6+        Ke7
4 Ke5        Ke8
5 Kf6        Kf8
6 e7+        Ke8
7 Bb5#

172VAR-2
1 Kg5        Ng4
2 Bxg4    fxg4
3 Kxg4    Kxe7
4 Kf5        Kf7
5 e6+        Ke7!
6 Ke5        Ke8!
7 Kf6        Kf8
8 e7+        Ke8
9 Ke6
Black draws by 
stalemate.

173VAR-1
1 Kh6        Bf7
2 Bxh7    Bxc4
Black's Bishop’s 
sacrifice for the remaining 
Pawn: draws.

173VAR-2
1 Kh6        Bf7
2 Bd3        Kg8
White takes the Rook 
Pawn with check.

173VAR-3
1 Kh6        Bf7
2 Bd3        Bg8
3 g5        Bf7
4 Bxh7    Bxc4
5 g6        Bg8
6 g7#

173VAR-4
1 Kh6        Bf7
2 Bd3        Be6
3 g5        Bg8
4 Bxh7!    Bxc4
5 g6
mate next move.

173VAR-5
1 Kh6        Bf7
2 Bd3        Be6
3 g5        Bg8
4 Bxh7!    Bxh7
5 g6        Bg8
6 g7#

173VAR-6
1 Kh6        Bf7
2 Bd3        Be6
3 g5        Bg8
4 Bxh7!    Bxh7
5 g6        Kg8
6 gxh7+    Kh8
7 Kg6        b5
8 c5        b4
9 c6        b3
10 c7        b2
11 c8=Q#

174VAR-1
1 g7        e1=Q
2 g8=Q    Qa5
3 Kd7+    Kb7
4 Qc8# 

174VAR-2
1 g7        e1=Q
2 g8=Q        Qh4+
3 Kc7+         Qd8
4 Qxd8#

174VAR-3
1 g7        e1=Q
2 g8=Q    Kb7
3 Qb3+    Kc6
4 Qb6+    Kd5
5 Qb5!    Ke6
6 Qe8+ 
winning the Queen.

174VAR-4
1 g7        e1=Q
2 g8=Q    Kb7
3 Qb3+    Kc6
4 Qb6+    Kd5
5 Qb5!    Ke4
6 Qe8+ 
winning the Queen.

174VAR-5
1 g7        e1=Q
2 g8=Q    Kb7
3 Qb3+    Kc6
4 Qb6+    Kd5
5 Qb5!    Qe6
6 Qb3+ 
winning the Queen.  

174VAR-6
1 g7        e1=Q
2 g8=Q    Kb7
3 Qb3+    Kc6
4 Qb6+    Kd5
5 Qb5!    Qe4
6 Qb7+ 
winning the Queen.

174VAR-7
1 g7        e1=Q
2 g8=Q    Kb7
3 Qb3+    Kc6
4 Qb6+    Kd5
5 Qb5!    Qh4+
6 Be7+ 
winning the Queen.  

174VAR-8
1 g7        e1=Q
2 g8=Q    Kb7
3 Qb3+    Kc6
4 Qb6+    Kd5
5 Qb5!    Qg3
6 Bf2+  
winning the Queen.  

174VAR-9
1 g7        e1=Q
2 g8=Q    Kb7
3 Qb3+    Kc6
4 Qb6+    Kd5
5 Qb5!    Qc3
6 Bb4+  
winning the Queen.  

174VAR-10
1 g7        e1=Q
2 g8=Q    Kb7
3 Qb3+    Kc6
4 Qb6+    Kd5
5 Qb5!    Qd2
6 Qd7+  
winning the Queen.  

174VAR-11
1 g7        e1=Q
2 g8=Q    Kb7
3 Qb3+    Kc6
4 Qb6+    Kd5
5 Qb5!    Qh1
6 Qb7+  
winning the Queen.  

174VAR-12
1 g7        e1=Q
2 g8=Q    Kb7
3 Qb3+    Kc6
4 Qb6+    Kd5
5 Qb5!    Qd1
6 Qd7+
winning the Queen.  

174VAR-13
1 g7        e1=Q
2 g8=Q    Kb7
3 Qb3+    Kc6
4 Qb6+    Kd5
5 Qb5!    Qc1
6 Be3+
winning the Queen.  

174VAR-14
1 g7        e1=Q
2 g8=Q    Kb7
3 Qb3+    Kc6
4 Qb6+    Kd5
5 Qb5!    Qa1
6 Bd4+    Kxd4
7 Qe5+
winning the Queen.

175VAR-1
1 Bc3    Kxa6
2 Bxb4
Black has been
stalemated.

175VAR-2
1 a7        Ka6
2 Kb8    Nc6+
3 Kc7        Nxa7
4 b4        Nc8
5 Kxc8    Ka7
6 Kc7        Ka6
7 Kc6        Ka7
8 Kxa6
and White wins.

175VAR-3
1 a7        Nc6
2 Kb7    Nxa7
3 Kxa7    Kb4
4 Kxb6    Kxb3
and Black draws.

176VAR-1
1 a5        Ne7
2 a6        Nd5
3 a7        Nb6
and the Pawns can 
go no further.

176VAR-2
1 Ba3        f5
2 d5!        cxd5
3 a5        Nf6
4 a6        Nd7 
5 a7        Nb6
where all Pawns
get pinned! 

176VAR-3
1 Ba3        f5
2 d5!        cxd5
3 a5        Nf6
4 a6        Nd7 
5 Bc5!        Nxc5
6 a7
White wins with
Black tempo loss. 

176VAR-4
1 Ba3        f5
2 d5!        cxd5
3 a5        Nf6
4 a6        Ne8
5 a7        Nc7
looks like a draw.

177VAR-1
1 Bh5!             gxh5
2 h7        Rc8
3 g6
Pawns cannot be stopped. 
One becomes a Queen.

177VAR-2
1 Bh5!             Kg3
2 Bxg6!        Rxg6
3 h7        Ra6+
4 Kb4        Ra8
5 g6
should be convincing.

177VAR-2
1 Bh5!             Kg3
2 Bxg6!        Kf4
3 h7        Rc3+
4 Kb4        Rh3
5 Bh5!    Rxh5
6 g6
and White wins.

178VAR-1
1 e8=Q        Rxe8
2 Bf8        Re2+
3 Kg3        Re3+
4 Kf4 

178VAR-2
1 e8=Q        Rxe8
2 Bf8        Re2+
3 Kg3        Re6
4 g8=Q    Rg6+
5 Qxg6    hxg6
lets Black get away
with a draw.

178VAR-3
1 e8=Q        Rxe8
2 Bf        Re2+
3 Kh3        Re3+
4 Kh4        Re4+
5 Kh5        Re5+
6 Kh6        Re1
7 g8=Q    Rh1+
8 Kg7        Rg1+
wins the Queen.

178VAR-4
1 e8=Q        Rxe8
2 Bf8        Re2+
3 Kh3        Re3+
4 Kh4        Re4+
5 Kh5        Re5+
6 Kh6        Re1
7 Bc5!        Rh1+
8 Kg5
Queen threatened with
check leaves Black no play, 
White wins.

178VAR-5
1 e8=Q    Rxe8
2 Bf8       Re2+
3 Kh3!    Re3+
4 Kh4    Re4+
5 Kh5    Re5+
6 Kh6    Re1
7 Bc5!    Re8
8 Kxh7    Rc8
9 Bf8       Ka8
10 g8=Q

178VAR-6
1 e8=Q        Rxe8
2 Bf8        Re2+
3 Kh3!        Re3+
4 Kh4        Re4+
5 Kh5        Re5+
6 Kh6        Re1
7 Bc5!        Re8
8 Kxh7        Rd8
9 Bf8        Rd7
pins the Pawn and 
draws, the Rook sacrificing
itself next move for the Pawn.

179VAR-1
1 a7        Bd5
2 c4        Bxc4
3 a8=Q+

180VAR-1
1 b3+!        Ka5
2 b4+        Ka4
3 c4        Be8
4 b5 
the passed Pawn is
unhindered to the 
Queening square.

180VAR-2
1 b3+!        Bxb3
2 c4        Bc2
3 a7        Be4+
4 Bf3
White wins.

180VAR-3
1 b3+!        Kxb3
2 c4        Be8
3 Bf3        e4
4 Bh5!        Bc6
5 Bf7        Kb4
6 Bd5
renders Black's 
Bishop impotent.

180VAR-3
1 b3+!        Kxb3
2 c4        Be8
3 Bf3        e4
4 Bh5!        Bc6
5 Bf7        Ka4
6 Be8
again renders Black's 
Bishop impotent.

182VAR-1
1 Kb7        h5
2 Kc6        h4
3 Kd5        Kg4
4 Ke4        h3
5 Rg8+    Kh4
6 Kf4        Kh5
7 Rh8+
winning the Pawn.

182VAR-2
1 Kb7        h5
2 Kc6        h4
3 Kd5        Kg4
4 Ke4        Kg3
5 Ke3        Kh3
6 Kf3
Black must abandon
 the Pawn.

182VAR-3
1 Kb7        h5
2 Kc6        h4
3 Kd5        Kg4
4 Ke4        Kg3
5 Ke3        Kg4
6 Kf2        h3
7 Rh7
again the Pawn falls.

182VAR-4
1 Kb7        h5
2 Kc6        h4
3 Kd5        Kg4
4 Ke4        Kg3
5 Ke3        h3
6 Rg8+        Kh4
7 Kf3        Kh5
8 Kg3
White wins the Pawn.

182VAR-5
1 Kb7        h5
2 Kc6        h4
3 Kd5        Kg4
4 Ke4        Kg3
5 Ke3        h3
6 Rg8+        Kh2
7 Kf2        Kh1
8 Kg3        h2
9 Rh8        Kg1
10 Rxh2
but actual move next
is faster since it 
forces mate.

182VAR-6
1 Kb7        h5
2 Kc6        h4
3 Kd5        Kg4
4 Ke4        Kg3
5 Ke3        h3
6 Rg8+        Kh2
7 Kf2        Kh1
8 Ra8        h2
9 Ra1#

183VAR-1
1 Kb7        g5
2 Kc6        g4
3 Kd5        Kf4
4 Kd4        g3
5 Rf8+               Kg4
6 Ke3               g2
7 Kf2
White wins.

183VAR-2
1 Kb7        g5
2 Kc6        g4
3 Kd5        Kf4
4 Kd4        g3
5 Rf8+        Kg4
6 Ke3        Kh3
7 Kf3        Kh2
8 Rh8+
and the Pawn is gone.

183VAR-3  
1 Kb7        g5
2 Kc6        g4
3 Kd5        Kf4
4 Kd4        Kf3
5 Kd3        g3
6 Rf8+    Kg4
7 Ke2        Kh3
8 Kf1        Kh2
9 Rh8#

183VAR-4  
1 Kb7        g5
2 Kc6        g4
3 Kd5        Kf4
4 Kd4        Kf3
5 Kd3        g3
6 Rf8+    Kg2
7 Ke2        Kh2
8 Rg8        g2
9 Kf2
and the Pawn 
is doomed. 

183VAR-5  
1 Kb7        g5
2 Kc6        g4
3 Kd5        Kf4
4 Kd4        Kf3
5 Kd3        g3
6 Rf8+    Kg2
7 Ke2        Kh2
8 Rg8        Kg2
9 Rg7    Kh2
10 Kf3
White wins the Pawn.

184VAR-1
1 Rh1     h5
in reply saves the 
Pawn and advances 
its career.

185VAR-1
1 Rg8+    Kf3
2 Rh8           Kg4
3 Kb2           h4
4 Kc2           h3
5 Kd2           Kg3
6 Ke2!           h2
7 Kf1
Black must give 
up the Pawn.

185VAR-2
1 Rg8+    Kf3
2 Rh8           Kg4
3 Kb2           h4
4 Kc2           h3
5 Kd2           Kg3
6 Ke2!           Kg2
7 Rg8+    Kh2
8 Kf2        Kh1
9 Rg1+    Kh2
10 Rg3    Kh1
11 Rxh3#

186VAR-1
1 Ke6        Kb3
2 Kd5    a3
3 Rb6+    Kc2
4 Kc4        a2
5 Ra6        Kb2
and White cannot
win and draws. 

186VAR-2
1 Ke6        Kb3
2 Kd5    a3
3 Kd4    a2
4 Kd3        Kb2
5 Rb6+    Kc1
6 Ra6        Kb2
and Black has 
a draw again.

187VAR-1
1 Rf8+    Ke4
2 Kf6    Kf4
3 Kg6+    Kg4
4 Rf5
with Black Pawn’s
loss and the game. 

187VAR-2
1 Rf8+    Ke4
2 Kf6    g4
3 Kg5        g3
4 Kh4        g2
5 Rg8    Kf3
6 Kh3
and the Pawn falls
on next move.

187VAR-3
1 Rf8+    Kg6
2 Rf6+    Kg7
3 Ke6        g4
4 Kf5        g3
5 Rg6+    g2
6 Rxg3
and White wins.

187VAR-4
1 Rf8+    Kg6
2 Rf6+    Kh5
3 Ke6         g4
4 Kf5        Kh4
5 Rh6+    Kg3
6 Rg6
again White wins 
the Pawn and game.

188VAR-1
1 Rc7+    Kh6
2 Rc2
Black loses his Knight
Pawn to prevent mate.

188VAR-2
1 Rc7+    Kg8
2 Rg7+    Kf8
3 Rb7
mate threat wins 
the Knight Pawn.

188VAR-3
1 Rc7+    Kg8
2 Rg7+    Kh8
3 Rb7        a3
4 Kg6        b1=Q+
5 Rxb1    a2
6 Rb8#

189VAR-1
1 Rg6!    Kd7
2 Rg4!    Ke6
3 Rxf4    Ke5
4 Rg4
White wins.

189VAR-2
1 Rg6!    Kd7
2 Rg4!    g2
3 Rxg2    Ke6
4 Rg5!    f3
5 Rg3
attacks and wins 
the Pawns.

189VAR-3
1 Rg6!    Kd7
2 Rg4!    g2
3 Rxg2    Ke6
4 Rg5!    Kf6
5 Rc5        Ke6
6 Kb7        Kd6?
7 Rf5
White wins.

189VAR-4
1 Rg6!    Kd7
2 Rg4!    g2
3 Rxg2    Ke6
4 Rg5!    Kf6
5 Rc5        Ke6
6 Kb7        f3
7 Rc3        f2
8 Rf3
Black loses Pawns
and game.

190VAR-1
1 Rd2+    Ka1
2 Kb3        g2
3 Rd1#

190VAR-2
1 Rd2+    Kb1
2 Kc3        g2
3 Rd1+    Ka2
4 Rg1    Ka3
5 Ra1# 

190VAR-3
1 Rd2+    Kb1
2 Kc3        g2
3 Rd1+    Ka2
4 Rg1    h2
5 Rxg2+    Ka3
6 Rxh2
a Pawn wipeout.

190VAR-4
1 Rd2+    Kb1
2 Kc3        h2
3 Rd1+    Ka2
4 Rh1    Ka3
5 Ra1#  

190VAR-5
1 Rd2+    Kb1
2 Kc3        h2
3 Rd1+    Ka2
4 Rh1    g2
5 Rxh2
last Pawn, being pinned, 
will fall next move.

190VAR-6
1 Rd2+    Kb1
2 Kc3        Kc1
3 Ra2        Kb1
4 Re2        g2
5 Re1+    Ka2
6 Rg1
White wins both 
Pawns.

190VAR-7
1 Rd2+    Kb1
2 Kc3        Kc1
3 Ra2        Kd1
4 Kd3        Ke1
5 Ke3    Kf1
6 Kf3
and capture of the
Pawns or simply
mate. 

190VAR-8
1 Rd2+    Kb1
2 Kc3        Kc1
3 Ra2        Kd1
4 Kd3        Ke1
5 Ke3    Kd1
6 Kf3    h2
7 Kxg3
and loses the Pawn
before or after the
Queen! 

190VAR-9
1 Rd2+    Kb1
2 Kc3        Kc1
3 Ra2        Kd1
4 Kd3        Ke1
5 Ke3    Kd1
6 Kf3    g2 
7 Kf2 
White wins.

190VAR-10
1 Rd2+    Kb1
2 Kc3        Kc1
3 Ra2        Kd1
4 Kd3        Kc1
5 Ke3        g2
6 Kf2
 is decisive. 

190VAR-11
1 Rd2+    Kb1
2 Kc3        Kc1
3 Ra2        Kd1
4 Kd3        Kc1
5 Ke3        Kb1
6 Re2        g2
7 Re1+    Kb2
8 Kf3
and Pawns are lost.

190VAR-12
1 Rd2+    Kb1
2 Kc3        Kc1
3 Ra2        Kd1
4 Kd3        Kc1
5 Ke3        h2
6 Ra1+    Kb2
7 Rh1        Kc3
8 Kf3
another way to loose
the Pawns.

191VAR-1
1 Rh2      a3
2 Ke5              a2
3 Kd4              Kb1
4 Kc3               a1=Q+
5 Kb3
Black must give up
his Queen to avoid mate.

191VAR-2
1 Rh2      a3
2 Ke5              a2
3 Kd4              Kb1
4 Kc3              a1=N
5 Rb2+    Kc1
6 Ra2                Kb1
7 Rxa6           Nc2
8 Re6!         Na1
9 Re2
to avoid mate, Black 
must move the Knight,
 and lose it.

191VAR-3
1 Rh2      a3
2 Ke5              a2
3 Kd4              Kb1
4 Kc3             a1=N
5 Rb2+    Kc1
6 Ra2                Kb1
7 Rxa6           Nc2
8 Re6!         Na3
9 Re1+    Ka2
10 Re2+    Kb1
11 Kb3
 threatens mate and 
 wins the Knight.

192VAR-1
1 Rg5+     Kh7
2 Rh5+
wins the Bishop.

192VAR-2
1 Rg5+     Kf8
2 Rh5                 Bg3
3 Rf5+    Kg7
4 Rg5+
wins the Bishop. 

192VAR-3
1 Rg5+     Kf8
2 Rh5                 Bg3
3 Rf5+    Ke8
4 Rg5
Black loses Bishop, 
to avoid mate.

192VAR-4
1 Rg5+     Kf8
2 Rh5                 Bf4
3 Rf5+
another Bishop loss.

192VAR-3
1 Rg5+    Kf8
2 Rh5        Bb8
3 Rh8+
winning the Bishop.

192VAR-4
1 Rg5+     Kf8
2 Rh5          Bc7
3 Kd7        Bg3
4 Rf5+     Kg8
5 Rg5+
again winning the 
Bishop.

193VAR-1
1 Kc3        a1=Q+
2 Kb3        Qa5!
scuttles the mate threat
and wins for Black.

193VAR-2
1 Re1+    Kb2
2 Ra1!    Kb3
3 Kd2        Kb2
4 Kd1        Kb3
5 Kc1    Ka3
6 Kc2        Kb4
7 Rxa2
White with a King-
Rook win.

193VAR-3
1 Re1+    Kb2
2 Ra1!    Kb3
3 Kd2        Kb2
4 Kd1        Kxa1
5 Kc2        a5
6 b6        a4
7 b7        a3
8 Kc3!        Kb1
9 b8=Q+    Ka1
10 Qd6    Kb1
11 Qd1#

194VAR-1
1 Kd4        d2
wins for Black.

194VAR-2
1 Kf4        d2
2 Rd6    e2
wins for Black.

194VAR-3
1 Kd6!    d2
2 Kc7        b6
3 Rxb6    d1=Q
4 Ra6#

195VAR-1
1 Rg5        b4+
2 Kc4    b3
3 Kc5
forces mate.

195VAR-2
1 Rg5        h3
2 Rg4+    b4+
3 Kc4        b3
4 Kc5#
discovers check 
and mate.

195VAR-3
1 Rg5        h3
2 Rg4+    b4+
3 Kc4        h2
4 Rg3        b3
5 Rxb3    g1=Q
6 Ra3#

196VAR-1
1 Rf7        Bg6
2 Rg7
wins a piece.

196VAR-2
1 Rf7        Be4
2 Rf4
pins and wins
the Bishop.

196VAR-3
1 Rf7        Be4
2 g5             Nh6
3 Ra7+    Kb4
4 Rb7+
 wins the Knight.
 
196VAR-4
1 Rf7        Be4
2 g5                Nh6
3 Ra7+    Kb3
4 Rg7
White wins it. 

196VAR-5
1 Rf7        Bb1
2 g5        Kb5
3 Rb7+
with Bishop loss.

196VAR-6
1 Rf7        Bb1
2 g5        Kb4
3 Rb7+
again Bishop goes.

196VAR-7
1 Rf7        Bb1
2 g5        Kb3
3 Rg7
winning the Knight.

196VAR-8
1 Rf7        Bb1
2 g5        Ka5
3 Rg7    Ba2
4 Ra7+
the skewered Black
King yields the Bishop. 

196VAR-9
1 Rf7        Bb1
2 g5        Ka3
3 Kc3     Ka4
4 Ra7+    Kb5
5 Rb7+
another skewer.

197VAR-1
1 e7        Rb8
2 Ra6+    Kb4
3 Rb6+     Rxb6
4 e8=Q
wins for White.

198VAR-1
1 h6        Rf8
2 h7        Kf6
3 Rf1+    Kg7
4 Rxf8
White wins.

198VAR-2
1 h6        Rf8
2 h7        Rh8
3 Kg5    Ke6
4 Kg6        Ke7
5 Kg7
Black has to give up 
his Rook for the Pawn.

199VAR-1
1 Ke4        Re8+
2 Kf5        Rd8
3 Ke5        Re8+
4 Kd6        Rd8+
5 Kc5        Rc8+
White makes no progress.
but gets the Pawn up a step, 
and enables White to carry 
out his idea.

199VAR-2
1 d5!        Rxd5+
2 Kc4
threatening mate and
attacking the Rook, 
is killing.

200VAR-1
1 Kb8        Rb2+
2 Kc8        Ra2
3 Rg6+    Kb5
4 Kb7
followed by Queening 
the Pawn.

200VAR-2
1 Kb8        Rb2+
2 Kc8        Ra2
3 Rg6+    Kd5
4 Kb7    Rb2+
5 Rb6    Ra2
6 Ra6
again the Pawn 
will Queen.

200VAR-3
1 Kb8        Rb2+
2 Kc8        Ra2
3 Rg6+    Kc5
4 Kb7        Rb2+
5 Kc7        Ra2
6 Rc6+    Kd5
7 Kb7        Rxa7+
brusquely forces a draw.

200VAR-4
1 Kb8    Rb2+
2 Kc8    Ra2
3 Rg6+    Kc5
4 Kb7    Rb2+
5 Kc7    Ra2
6 Rg5+    Kc4
7 Kb7     Rb2+
8 Kc6     Ra2
9 Rg4+    Kc3
10 Kb6    Rb2+
11 Ka5       Ra2
12 Ra4         Rb2
13 a8=Q

201VAR-1
1 Kf4        Kf2
2 Ke4        Ra4+
3 Kd3    Ra3+
4 Kc2        Ra2+
5 Kb3        Ra1
6 Rf8+
and White wins.

201VAR-2
1 Kf4        Kf2
2 Ke4        Ke2
3 Kd4    Kd2
4 Kc5        Rc1+
5 Kb4        Rb1+
6 Ka3        Ra1+
7 Kb2        Ra3?
8 Rd8+
White wins Rook
and game.

202VAR-1
1 Rc8        Kd7
2 Rb8    Rh1
3 Kb7        Rb1+
4 Ka6        Ra1+
5 Kb6        Rb1+
6 Kc5
White wins.

202VAR-2
1 Rc8        Kd6
2 Rb8        Rh1
3 Kb7        Rb1+
4 Kc8        Rc1+
5 Kd8        Rh1
6 Ke8        Rh8+
7 Kf7    Rh7+
wins the Pawn.

202VAR-3
1 Rc8        Kd6
2 Rb8        Rh1
3 Kb7        Rb1+
4 Kc8        Rc1+
5 Kd8        Rh1
6 Rb6+    Ke5
7 Kc7    Rh7+
8 Kb8        Rh8+
9 Kb7        Rh7+
10 Ka6     Rh8
11 Rb8
and White wins.

202VAR-4
1 Rc8        Kd6
2 Rb8        Rh1
3 Kb7        Rb1+
4 Kc8        Rc1+
5 Kd8        Rh1
6 Rb6+    Kc5
7 Re6        Ra1
8 Re7        Kb6
draws. 

202VAR-5
1 Rc8        Kd6
2 Rb8        Rh1
3 Kb7        Rb1+
4 Kc8        Rc1+
5 Kd8        Rh1
6 Rb6+    Kc5
7 Ra6        Rh8+
8 Ke7        Rh7+
9 Kf8         Rh8+
10 Kg7     Ra8
11 Kf7    Kb5
12 Ra1    Kb6
Black draws by 
taking the Pawn.

202VAR-6
1 Rc8        Kd6
2 Rb8        Rh1
3 Kb7        Rb1+
4 Kc8        Rc1+
5 Kd8        Rh1
6 Rb6+    Kc5
7 Rc6+!    Kxb5
8 a8=Q+
winning the Rook
and the game. 

202VAR-7
1 Rc8        Kd6
2 Rb8        Rh1
3 Kb7        Rb1+
4 Kc8        Rc1+
5 Kd8        Rh1
6 Rb6+    Kc5
7 Rc6+!    Kd5
8 Ra6        Rh7+
9 Kc7        Rh7+
10 Kb6        Rh6+
11 Kb5
White wins.

203VAR-1
1 a7        Rb1+
2 Ka6        Ra1+
3 Kb6        Rb1+
4 Kc5        Rc1+
5 Kb4    Rb1+
6 Kc3        Rc1+
7 Kb2        Rc7
8 Rh8        Rxa7
9 Rh7+
again wins the Rook
and the game. 

204VAR-1
1 Kb5        Rb8+
2 Kc6        Rc8+
3 Kb7        Rc1
4 a6        Ra1
5 Rb3
King is shielded
from checks--Pawn
marches up without 
hindrance.

204VAR-2
1 Kb5        Rb8+
2 Kc6        Rc8+
3 Kb7        Rc1
4 a6        Rb1+
5 Kc6        Rc1+
6 Kb5        Rb1+
7 Ka4        Ra1+
8 Ra3    Rh1
9 a7        Rh8
10 Kb5        Ra8
11 Kb6
White wins easily.

204VAR-3
1 Kb5        Rb8+
2 Kc6        Rc8+
3 Kb7        Rc1
4 a6        Rb1+
5 Kc6        Rc1+
6 Kb5        Rb1+
7 Ka4        Rb8
8 Ka5        Rb1
9 Ra3
is decisive for Pawn
Queening protectiton.

205VAR-1
1 Kb4        Re7
2 Rxe7+    Kxe7
3 Kb5        Kd7
4 Kb6        Kc8
5 a5         Kb8
Black has an easy draw.

205VAR-2
1 Kb4        Re7
2 Rxe7+    Kxe7
3 Kb5        Kd7
4 Kb6        Kc8
5 Ka7    Kc7    
again Black has an 
easy draw with Kings
opposition.

205VAR-3
1 Kb4        Re7
2 Ra3        Ke8
3 a5        Kd8
4 a6        Ra7
5 Kb5        Kc8
6 Rh3        Rc7
7 Rh8+    Kd7
8 a7        Rxa7
9 Rh7+
winning the Rook.

205VAR-4
1 Kb4        Re7
2 Ra3        Ke8
3 a5        Kd8
4 a6        Ra7
5 Kb5        Kc8
6 Rh3        Rd7
7 Rh8+    Kc7
8 a7
the Pawn will Queen.

205VAR-5
1 Kb4        Re7
2 Ra3        Ke8
3 a5        Kd8
4 a6        Ra7
5 Kb5        Kc8
6 Rh3        Rc7
7 Rh8+    Kd7
8 a7        Rxa7
9 Rh7+
winning the Rook.

205VAR-6
1 Kb4        Re7
2 Ra3        Ke8
3 a5        Kd8
4 a6        Ra7
5 Kb5        Kc8
6 Rh3        Rd7
7 Rh8+    Kc7
8 a7
the Pawn will Queen.

205VAR-7
1 Kb4        Re7
2 Ra3        Ke8
3 a5        Kd8
4 a6        Ra7
5 Kb5        Kc8
6 Rh3        Kb8
7 Kb6        Rb7+!
8 axb7
stalemates Black. 

205VAR-8
1 Kb4        Re7
2 Ra3        Ke8
3 a5        Kd8
4 a6        Ra7
5 Kb5        Kc8
6 Rh3        Kb8
7 Kb6        Rb7+!
8 Kc5        Rb1
and White has been 
swindled out of a win.

205VAR-9
1 Kb4        Re7
2 Ra3        Ke8
3 a5        Kd8
4 a6        Ra7
5 Kb5        Kc8
6 Rh3        Kb8
7 Rh8+    Kc7
8 Rg8     Kd7
loses a Rook by 9 Rg7+

206VAR-1
1 Rh7+    Kg5
2 g7             Kg6
3 Kh8           Ra8
followed by
Pawn Queening.

206VAR-2
1 Rh7+    Kg5
2 g7             Ra8+
3 Kf7        Ra7+
4 Ke6        Ra6+
5 Ke5        Rg6
6. Rh8    Rxg7
draws with a
Rook move.

206VAR-3
1 Rh7+    Kg5
2 g7             Ra8+
3 Kf7        Ra7+
4 Ke6        Ra6+
5 Ke5        Rg6
6 Kd5    Kf6
and the Pawn will 
fall -- a draw.

206VAR-4
1 Rh7+    Kg5
2 g7             Ra8+
3 Kf7        Ra7+
4 Ke6        Ra6+
5 Ke5        Ra8
6 Rh8
heading off the Pawn
was ineffective.

206VAR-5
1 Rh7+    Kg5
2 g7             Ra8+
3 Kf7        Ra7+
4 Ke6        Ra6+
5 Ke5        Ra8
6 Ke5!    Kg4
7 Rh1        Rg5+
8 Ke4    Rxg7
9 Rg1+
 wins the Rook. 

206VAR-6
1 Rh7+    Kg5
2 g7        Ra8+
3 Kf7        Ra7+
4 Ke6        Ra6+
5 Kd5!    Rg6
6 Ke5!    Kg4
7 Rh1        Rg5+
8 Ke4    Kg3
9 Rg1+    Kh4
10 Rxg5    Kxg5
11 g8=Q+
and White wins.

207VAR-1
1 Kb2        Rb8
2 Rb3        Kf7
3 b7         Ke7
4 Ka3        Kd7
5 Ka4        Kc7
and the Pawn is lost.

207VAR-2
1 Kb2        Rb8
2 Rb3        Kf7
3 Ka3!    Ke7
4 Ka4        Kd7
5 Kb5        Rb7 
6 Rc3        Rb8
7 Rd3+    Kc8
8 Kc6    Rb7
9 Rd8+    Kxd8
10 Kxb7    Kd7
11 Ka7 
and White wins.

207VAR-3
1 Kb2        Rb8
2 Rb3        Kf7
3 Ka3!    Ke7
4 Ka4        Kd7
5 Kb5        Rb7 
6 Rc3        Kd8
7 Ka6        Rb8
8 Ka7        Rc8
9 Rxc8+
White wins again.

207VAR-4
1 Kb2        Rb8
2 Rb3        Kf7
3 Ka3!    Ke7
4 Ka4        Kd7
5 Kb5        Rb7
6 Ka6        Kc8
7 Rc3+    Kb8
8 Rh3    Kc8
9 Rh8+
King deflection
wins the Rook. 

207VAR-5
1 Kb2        Rb8
2 Rb3        Kf7
3 Ka3!    Ke7
4 Ka4        Kd7
5 Kb5        Rb7
6 Ka6        Kc8
7 Rc3+    Kb8
8 Rh3    Ka8
9 Rh8+    Rb8
10 b7# 
sweet.

207VAR-6
1 Kb2        Rb8
2 Rb3        Kf7
3 Ka3!    Ke7
4 Ka4        Kd7
5 Kb5        Kc8
6 Rc3+           Kd7
7 Rd3+    Ke7
8 Kc6
with easy White win. 

207VAR-7
1 Kb2        Rb8
2 Rb3        Kf7
3 Ka3!    Ke7
4 Ka4        Kd7
5 Kb5        Kc8
6 Rc3+            Kd7
7 Rd3+    Kc8
8 Kc6        Rb7
9 Rd8+    Kxd8
10 Kxb7
the Pawn will Queen.

208VAR-1
1 Rd4!    Rd8
2 Rxd8    Kxd8
3 Ka4!    Kc8
4 Kb5        Kb7
Black, having the 
opposition, draws. 

208VAR-2
1 Rd4!    Rd8
2 Rxd8    Kxd8
3 Ka4!    Kc8
4 Ka5        Kc7
5 Ka6        Kb8
6 Kb6
White wins.

208VAR-3
1 Rd4!    Ke6
2 Kc4!        Ke5
3 Rd5+    Ke6
4 b5        Rc8+
5 Rc5        Rxc5+
6 Kxc5    Kd7
7 Kb6    Kc8
8 Ka7 
White wins.

208VAR-4
1 Rd4!    Ke6
2 Kc4!        Ke5
3 Rd5+    Ke6
4 b5        Rc8+
5 Rc5        Kd7
6 Rxc8    Kxc8
7 Kc5        Kc7
Black has a draw.

208VAR-5
1 Rd4!    Ke6
2 Kc4!        Ke5
3 Rd5+    Ke6
4 b5        Rc8+
5 Rc5        Kd7
6 b6        Rxc5+
7 Kxc5    Kd8
8 Kd6!    Kc8
9 Kc6
and White wins.

209VAR-1
1 Kc3        Rc8+
2 Kd4        Rb8
3 Kc4        Rc8+
4 Kd5        Rd8+
5 Kc6        Rc8+
6 Kb7        Rc3
7 Rb1        Ke6
8 b4        Kc6
9 b5
White wins.

209VAR-2
1 Kc3        Rc8+
2 Kd4        Rb8
3 Kc4        Rc8+
4 Kd5        Rb8
5 Rb1        Rb5+
6 Kc6        Rb4
7 Kc5        Rb8
8 b4        Ke6
9 Kc6        Rc8+
10 Kb7
White has no 
trouble winning.

210VAR-1   
1 d6        Rh6
2 d7        Rc6+
3 Kb4        Rd6
Black wins the
Pawn.

210VAR-2
1 d6        Rh6
2 Rd2        Rh8
3 d7        Rd8
4 Kc4        Kb6
and the Pawn 
will fall. 

210VAR-3
1 d6        Rh6
2 Rd2     Rh8
3 Kd4        Kb7     
4 Rc2        Rh4+
White cannot escape
perpetual check or 
loss of the Pawn. 

210VAR-4
1 Kd3!    Rg4
2 Ke3        Rh4
3 d6        Rh6
4 d7        Re6+
5 Kf3        Rd6
will remove the Pawn.

210VAR-5
1 Kd3!    Rg4
2 Ke3        Rh4
3 d6        Rh6
4 Rd2!    Rh8
5 d7        Rd8
6 Ke4     Kb7
7 Ke5     Kc7
8 Ke6     Kc6
9 Rc2+    Kb7
10 Ke7
Black must give up
his Rook for the Pawn
White wins.

210VAR-6
1 Kd3!    Rg4
2 Ke3        Rh4
3 d6        Rh6
4 Rd2!    Rh8
5 d7        Rd8
6 Ke4     Kb7
7 Ke5     Kc7
8 Ke6     Rh8
9 Rc2+    Kd8
10 Rc8#

210VAR-7
1 Kd3!    Rg4
2 Ke3        Rh4
3 d6        Rh6
4 Rd2!    Rh8
5 d7        Rd8
6 Ke4     Kb7
7 Ke5     Kc7
8 Ke6     Rh8
9 Rc2+    Kb7
10 Rh2!    Rxh2
allows White to Queen 
his Pawn and win. 

211VAR-1
1 Ke7        Re2+
2 Kd6        Rd2+
3 Kc6        Rc2+
4 Kb5     Rb2+
5 Kc4        Rd2
Black wins the Pawn 
and draws.

211VAR-2
1 Ra1        Kf7
2 Ra8        Ke6
3 Ke8        Rh2
4 Ra6+    Kf5
5 d8=Q
and White wins.

211VAR-3
1 Ra1        Kf7
2 Ra8        Rc1
3 Rc8        Rd1
4 Kc7        Rc1+
5 Kb6        Rb1+
6 Kc5
White zigzags down the 
board until Black runs
out of checks, then
Queens his Pawn and 
wins.

211VAR-4
1 Rf4!        Rc1
2 Ke7         Re1+
3 Kd6        Rd1+
4 Ke6        Rd2
5 Rf5           Re2
6 Re5
shutting off Black's 
Rook, insures the 
advance of the Pawn.

212VAR-1
1 Ke6        Ra8
2 Rg1+    Kh7
3 Kf7
threatening mate as 
well as Queening the
Pawn, White wins.

212VAR-2
1 Ke6        Ra8
2 Kf5        Kf7
3 Re1    Ra1
4 Re2        Re1
5 e8=Q+ 
and White wins. 

212VAR-3
1 Ke6        Ra8
2 Ke5+!    Ra5+
3 Kf6        Ra6+
4 Kg5        Re6
5 Rf8+    Kg7
6 e8=Q+
White wins.

213VAR-1
1 Kd8        Kf6
2 e7!         Rxe7
3 Rf1+        Ke6
4 Re1+
costs Black his Rook.

213VAR-2
1 Kd8        Kf6
2 e7!        Rb8+
3 Kc7        Re8
4 Kd6        Rb8
5 Rf1+    Kg7
6 Kc7        Ra8
7 Ra1!    Rh8
8 Kd7    Kf7
9 Rf1+    Kg7
10 e8=Q
 and White wins.

214VAR-1
1 Kd3!    Kh5
2 e6!        Ra6
3 Ke4!    Rxe6+
4 Kf5
White's attack on Rook 
and a threat of mate on
the move forces Black
to give up his Rook.

214VAR-2
1 Kd3!    Kh5
2 e6!    Kh6
3 e7        Ra8
4 Ke4        Re8
5 Ke5!    Rxe7+
6 Kf6
again Black must lose
his Rook or be 
instantly mated.

214VAR-3
1 Kd3!    Kh5
2 e6!        Kh4
3 Rg8        Ra7 
otherwise the Pawn 
simply moves on 
to e7 and e8.

214VAR-4
1 Kd3!    Kh5
2 e6!        Kh4
3 Rg8        Ra7
4 Ke4        Kh5
5 Kf5        Kh6
6 Kf6        Kh7
7 Rd8    Ra6
8 Kf7
and the rest plays itself.

214VAR-5
1 Kd3!    Ra6 
2 Ke4
3 Kf5
Black could offer
no resistance, his
King being confined
to the Rook file.

214VAR-6
1 Kd3!    Rb4
2 e6!        Rb6
3 Re1        Rb8
4 e7        Re8
5 Kd4        Kg7
6 Kd5        Kf7
7 Kd6        
8 Rf1+    Kg7
9 Kd7        Ra8
10 e8=Q
Black will have to 
give up his Rook.

214VAR-7
1 Kd3!    Rb4
2 e6!        Rb6
3 Re1        Rb8
4 e7        Re8
5 Kd4        Kg7
6 Kd5        Kf7
7 Kd6        Rh8
8 Kd7    Re8
9 Rf1+
and Black must 
abandon his Rook.

214VAR-8
1 Kd3!    Ra6
2 e6!        Rb6
3 Re1        Rb8
4 e7        Re8
5 Kd4        Kg7
6 Kd5        Kf7
7 Kd6        Ra8
8 Rf1+        Ke8
the penalty is 9 Rf8#.
Gorgeous!

214VAR-9
1 Kd3!    Rb4 
2 e6!        Rb6
3 Re1        Rb8
4 e7        Re8
5 Kd4        Kg7
6 Kd5        Kf7
7 Kd6        Ra8
8 Rf1+        Kg7
9 Kd7
King cannot escape 
without loss of the 
Pawn. 

214VAR-10
1 Kd3!    Rb4 
2 e6!        Rb6
3 Re1        Rb8
4 e7        Re8
5 Kd4        Kg7
6 Kd5        Kf7
7 Kd6        Ra8
8 Rf1+        Kg7
9 Kd7    Ra7+
10 Kd6        Ra6+
11 Kd5    Ra5+
12 Kc6    Ra6+
13 Kb7        Re6
wins the Pawn.

214VAR-11
1 Kd3!    Rb4 
2 e6!        Rb6
3 Re1        Rb8
4 e7        Re8
5 Kd4        Kg7
6 Kd5        Kf7
7 Kd6        Ra8
8 Rf1+    Kg7
9 Kd7        Ra7+
10 Kd6        Ra6+
11 Kd5        Ra5+
12 Kc6        Ra6+
13 Kb5        Re6
wins the Pawn.

214VAR-12
1 Kd3!    Rb4 
2 e6!        Rb6
3 Re1        Rb8
4 e7        Re8
5 Kd4        Kg7
6 Kd5        Kf7
7 Kd6        Ra8
8 Rf1+        Kg7
9 Ra1!    Rxa1
10 e8=Q
wins for White.

214VAR-13
1 Kd3!    Rb4 
2 e6!        Rb6
3 Re1        Rb8
4 e7        Re8
5 Kd4        Kg7
6 Kd5        Kf7
7 Kd6        Ra8
8 Rf1+        Kg7
9 Ra1!    Rb8
10 Kc7    Re8
11 Kd7    Rb8
12 e8=Q

215VAR-1
1 Kd3!    Rd8+
2 Kc4        Rc8+
3 Kd5        Rd8+
4 Ke5        Re8+
5 Kf6!    Kh5
6 Re1
now the Pawn can
advance.

216VAR-1
1 Rd2+    Ke7
2 Kc7     Rc3+
King now has to return.

216VAR-2
1 Rd2+    Ke7
2 Rd6!    Kxd6
3 Kc8    Rc3+
4 Kd8    Rh3
5 b8
Pawn’s promoting with 
check wins for White.

218VAR-1
1 Ra8        Rxh3+
2 Kb4    Kb7
3 Rb8+     Kxc7 
after White Rook has 
scurried away Black 
can draw against the 
remaining Pawn.

218VAR-2
1 Rd8        Kxc7
2 Rd3
White has an easy win.

218VAR-3
1 Rd8        Rxh3+
2 Rd3!    Rh8
3 Rc3
leaves no hope. 

218VAR-4
1 Rd8        Rxh3+
2 Rd3!    Rxd3+
3 Kc4     Rd1
4 c8=Q     Rc1
(with or without 
check) draws.

218VAR-5
1 Rd8        Rxh3+
2 Rd3!    Rxd3+
3 Kc2!    Rd6!
4 c8=Q    Rc6+
5 Qxc6+    Kxc6
6 Kb3        Kb5
Black, having the 
opposition, draws. 

219VAR-1
1 Rh3+    Kg5
2 Rg3+    Kf5
3 Rxg6    Kxg6
4 a7
and White wins.

219VAR-2
1 Rh3+    Kg7
2 Rg3!    b2
3 Rxg6    Kxg6
4 Kxb2
If Black refuses the Rook--
Rooks are exchanged
with White Queening
his Pawn.

219VAR-3
1 Rh3+    Kg7    
2 Rg3!    Rxg3
3 a7        Rg1
4 Kb2        Rg2+
5 Kxb3    Rg3+
6 Kb4        Ra3
with Rook ruin and
Pawn Queening. 

219VAR-4
1 Rh3+    Kg7
2 Rg3!    Rxg3
3 a7        Rg1
4 Kb2        Rg2+
5 Kxb3    Rg3+
6 Ka4    Rg1
and 7. . . Ra1
with or without check.

219VAR-5
1 Rh3+    Kg7
2 Rg3!    Rxg3
3 a7        Rg1
4 Kb2        Rg2+
5 Kxb3    Rg3+
6 Kb4        Ra3
With Rook ruin and
Pawn Queening. 

219VAR-6
1 Rh3+    Kg7    
2 Rg3!    Rxg3
3 a7        Rg1
4 Kb2        Rg2+
5 Kxb3    Rg3+
6 Ka4    Rg1
7 Ka5    Ra1
Black draws with
or without check.

220VAR-1
1 Rg1        Rg3+
2 Kf2        Kh5
3 g7        Kh6
4 g8=Q
White wins.

220VAR-2
1 Rg1        Ra8
2 g7        Kh5
3 Kf3        Rg8
4 Kxf4        Kh6
5 Kf5        Rxg7
6 Rh1#

221VAR-1
1 Rxf5     Ra7+
leads to a drawing
 position.

221VAR-2
1 exf5    Ra7+
another drawing position.

221VAR-3
1 Rg1+    Kh7
2 e5!    Ra7+
3 Kf6    Ra6+     
4 e6
White  wins easily, with
Black's King cut off. 
from the action.

222VAR-1
1 Kg5!    Rh8
2 Ra3+    Kf2
3 Kxg4    Rc8
4 Rc3        Ke2
5 Kf4        Kd2
6 Rc5    Kd3
7 Ke5
and White wins.

222VAR-2
1 Kg5!    Rh1
2 Rc2        Kf3
3 Rc3+    Ke4
4 c7        Rh8
5 c8=Q    Rxc8
6 Rxc8    Pg3
7 Kh4        g2
8 Rg8        Kf3
9 Kh3
White wins again.

223VAR-1
1 b7        Ra7
2 Re1+        Kd8
3 Re7!    Kxe7
4 b8=Q
White wins.

223VAR-2
1 b7        Ra7
2 Re1+        Kd8
3 Re7!    Kxe7
4 b8=Q    Ra4
5 Qc7     Ke8
6 Qc6+
Black Rook and 
game lost.

224VAR-1
1 Rb8        Rd7+
2 Ke8        Rxa7
allows Black to draw.

224VAR-2
1 Ke7!    Rh8
2 Rb8
Black losing immediately.

224VAR-3
1 Ke7!    Ra8
2 Kd7!    Rxa7+
3 Kc6        Ra8
4 Ra2#
pushes Rook in all
directions for 
White win.

224VAR-4
1 Ke7!    Ra8
2 Kd7!    Rf8
3 Rf2!     Rg8
4 Kc7           Ra8
5 Ra2+         Kb5
6 Kb7
extracting the Rook
from  a8  and Queening
for White win. 

224VAR-5
1 Ke7!    Ra8
2 Kd7!    Rf8
3 Rf2!     Rh8
4 Kc7       Rh7
5 Ra2+     Kb5
6 a8=Q
a big White win. 

224VAR-6
1 Ke7!    Ra8
2 Kd7!    Rf8
3 Rf2!     Rg8
4 Kc7        Ra8
5 Ra2+    Kb5
6 Kb7
and White wins.

226VAR-1
1 Rf8        Re5
2 Rf6+        Kh5
3 Kg7        Rg5+
4 Kf7
Pawn moves on to 
become a Queen and win.

226VAR-2
1 Rf8        Rg5
2 Rf6+        Kh5
3 Rf5!        Rxf5
4 Kg7        Rg5+
5 Kf6        Rg6+
6 Kf7        Rh6
 wins for  Black.

227VAR-1
1 Rc2+        Kd5
2 b7        Rh2+
3 Kd3        Rh3
4 Kd2        Rb3
White’s Pawn is lost
and most probable
draw. 

227VAR-2
1 Rc2+    Kd5
2 b7        Rh2+
3 Kd1        Rh1+
4 Kd2        Rb1
again Black draws 
easily, his Rook in 
another way sneaks
behind the Pawn.

227VAR-3
1 b7        Rh2+
2 Ke3        Rh8
3 Rc2+    Kd5
4 Rc8
and White wins.

227VAR-4
1 b7        Rh2+
2 Ke3        Rh3+
3 Kf4        Rb3
4 Rc2+    Kd5
and Black has 
a drawn position.

228VAR-1
1 Rxg7        Rh6
 in reply White can 
make no progress.

228VAR-2
1 Rg1+        Kh2
2 Rg2+        Kh3
3 Kg1
4 Rh2+
wins for White.

228VAR-3
1 Rg1+        Kh2
2 Rg2+        Kh1
3 Kg3!        Rxh7
4 Rd2
with a dire mate 
threat would be 
the penalty.

228VAR-4
1 Rg1+        Kh2
2 Rg2+        Kh1
3 Kg3!        Rg5+
4 Kf3        Rxg2
5 h8=Q+
White winning. 

228VAR-5
1 Rg1+        Kh2
2 Rg2+        Kh1
3 Kg3!        Rg5+
4 Kf3        Rh5
5 Rxg7    Kh2
6 Kf4        Rh6
7 Kg5        Rh3
8 Kg6        Rh4
9 Kf7        Rh6
10 Kf8
and White wins.

228VAR-6
1 Rg1+        Kh2
2 Rg2+        Kh1
3 Kg3!        Rg5+
4 Kf3        Rf5+
5 Kg4        Rf8
6 Kg3        Rh8
7 Re2
and it's all over after
8 Re1#.

229VAR-1
1 e7        Re5
2 Kxe5        e2
3 Rf3+        Kd2
4 Re3
Rook gets behind 
the Pawn. 

229VAR-2
1 e7        Re5
2 Kxe5        e2
3 Rf3+        Kd2 
4 Rf2        Kd1
5 Rxe2
and White wins.

229VAR-3
1 e7        Re5
2 Rf8        e2
White captures the 
Rook with check, 
in response.

229VAR-4
1 e7        Rc8
2 Rf8        Rc6+
3 Kf5        Re6!
4 Kxe6        e2
5 Rf3+        Kd2
6 Rf2
pins the Pawn 
and wins.

230VAR-1
1 g7        Rd8
2 Re6+        Kd4
3 Rd6+    Rxd6
4 g8=Q
White wins.

230VAR-2
1 g7        Rd8
2 Re6+        Kf4
3 Rf6+        Kg3
4 Rf8        Rd1+
in reply would be fatal.

230VAR-3
1 g7        Rd8
2 Re6+        Kf4
3 Rf6+        Kg3
4 Rg6+        Kh3
5 g8=Q    Rxg8
6 Rxg8
and Black is 
stalemated.

230VAR-3
1 g7        Rd8
2 Re6+        Kf4
3 Rf6+        Kg3
4 Rg6+        Kh3
5 Kg1        Rg8
6 . . .         Rxg7
7 Rxg7
and draws by 
stalemate.

230VAR-4
1 g7        Rd8
2 Re6+        Kf4
3 Rf6+        Kg3
4 Rg6+        Kh3
5 Kg1        Rg8
6 Kf2        Kh2
7 Rg4!        h3
8 Rg3        Kh1
9 Rxh3#

230VAR-5
1 g7        Rd8
2 Re6+        Kf4
3 Rf6+        Kg3
4 Rg6+        Kh3
5 Kg1        Rg8
6 Kf2        Kh2
7 Rg4!        Kh3
8 Kf3        Kh2
9 Rxh4+    Kg1
10 Rh7        Kf1
11 Rh1#

231VAR-1
1 Kc3        Ka2
2 g7        Rg2
3 Rb2+     Rxb2
4 g8=Q+
and White wins.

232VAR-1
1 Kb8        Rc6
2 Rf6        Rxf6
3 gxf6    c4
4 f7        c3
5 f8=Q    c2
6 Qc5
and White wins.

232VAR-2
1 Kb8        Rc6
2 Rf6        c4
3 Rxc6    dxc6
4 g6        c3
5 g7        c2
6 g8=Q    c1=Q
the position is a draw.

232VAR-3
1 Kb8        Rc6
2 Rf6        c4
3 Rh6+        Rxh6
4 gxh6    c3
5 h7        c2
6 h8=Q
the Pawn Queens with 
check, wins for White. 

233VAR-1
1 Rb7+        Kc3
2 Rxc7+       Kb2
3 Rxc1        Kxc1
Rooks are exchanged 
then pushing the Pawn, 
wins on the spot.

233VAR-2
1 Rb7+        Ka3
2 g7        Rg1
3 Rxc7        Kb2
4 Re7         Kc1
5 Re1+    Rxe1
6 g8=Q
wins for White.

233VAR-3
1 Rb7+        Ka3
2 g7        Rg1
3 Rxc7        Kb2
4 Kh4        f5
5 Kh5        f4
6 Kh6        f3
7 Rf7        Rg3
8 Kh7        Rh3+
9 Kg8        Kc4
10 Kf8
again White wins.

233VAR-4
1 Rb7+        Ka3
2 g7        Rg1
3 Rxc7        Kb2
4 Kh4        f5
5Re7        Kc3
6 Re3+    Kd4
7 Rg3
and White wins.

233VAR-5
1 Rb7+        Ka3
2 g7        Rg1
3 Rxc7        Kb4!
4 Kh4        Rg5
5 Re7        Ka5
6 Re5+!    fxe5
7 Kxg5
White wins.

233VAR-6
1 Rb7+        Ka3
2 g7        Rg1
3 Rxc7        Kb4!
4 Kh4        Rg5
5 Re7        Kc5
6 Re5+!    Kd6
7 Rxg5    fxg5+
8 Kxg5
game is over with the 
solo Pawn Queening.

234VAR-1
1 Kf7        f3    
2 Rxh6#

234VAR-2
1 Kf7        Kh7
2 Rg7+    Kh8
3 Kg6        f3
4 Re7
forces mate.

234VAR-3
1 Kf7        Rxh5
2 Rg8+        Kh7
3 Rg7+        Kh8
4 Kg6        g4
5 Rd7
followed by mate.

234VAR-4
1 Kf7        Rxh5
2 Rg8+        Kh7
3 Rg7+        Kh8
4 Kg6        g4
5 Ra7        Rg5+
6 Kxh6        Rg8
7 Rh7#
the penalty being mate. 

234VAR-5
1 Kf7        Rxh5
2 Rg8+        Kh7
3 Rg7+        Kh8
4 Kg6        g4
5 Ra7        Rg5+
6 Kxh6        g3
7 Ra8+
 and mate next.

234VAR-6
1 Kf7        Rxh5
2 Rg8+        Kh7
3 Rg7+        Kh8
4 Kg6        g4
5 Ra7        Rg5+
6 Kxh6        g3
7 Kxg5        Kg8
8 Kxf4
the last Pawn will
also fall.

235VAR-1
1 Kc7        a6
2 Kb6        Rb5+
3 Kxa6        Rb7
White has nothing to fear.

235VAR-2
1 e6!        Ra6
2 exf7        Rxh6
3 f8=Q+
White wins.

235VAR-3
1 e6!        fxe6
2 Kc6        Ra6+
3 Kc7
forces a quick mate.

235VAR-4
1 e6!        fxe6
2 Kc6        a6
3 Rh8+        Ka7
4 Rh7+        Kb8
5 Kb6        Rb5+
6 Kxa6
and the Rook is trapped.

236VAR-1
1 g7        Rg6
2 Ra1        Kb8
3 c7+        Kc8
4 g8=Q+    Rxg8
5 Ra8+
with a great snooker
White wins the Rook.

236VAR-2
1 g7        Rg6
2 Ra1        Kb8
3 c7+        Kb7
4 c8=Q+    Kxc8
5 g8=Q+    Rxg8
6 Ra8+
 again White snookers 
the Rook.

237VAR-1
1 b7        Kc6
makes no headway

237VAR-2
1 Rb7        Kc6
after the exchange
 of Pawns, Black has 
no trouble drawing.

237VAR-3
1 Rd8+!    Kxd8
2 b7        Rb4!
3 Kxb4        c5+
4 Kxc5    Kc7
5 Kb5        Kxb7
Black has nothing to 
fear from the Rook Pawn.

237VAR-4
1 Rd8+!    Kxd8
2 b7        Rb4!
3 Kxb4        c5+
4 Kb5!        Kc7
5 Ka6
6 Ka7
followed by Queening 
with check.

238VAR-1
1 d5        Ra8
2 d6        Kf7
3 d7        Ke7
4 d8=Q+    Rxd8
5 Rxd8    Kxd8
6 a6        Kc7
Black overtakes the
Rook Pawn and draws.

238VAR-2
1 d5        Ra8
2 d6        Kf7
3 d7        Ke7
4 a6        Kd8
5 Rd6        Kc7
6 d8=Q+    Rxd8
7 Rxd8        Kxd8
8 a7
and White wins.

238VAR-3
1 d5        Ra8
2 d6        Kf7
3 d7        Ke7
4 a6        Kd8
5 Rd6        Ra7
6 Rxg6    Rxa6
7 Rg8+    Ke7
8 Rg7+
White Pawn Queens
and wins.

238VAR-4
1 d5        Ra8
2 d6        Kf7
3 d7        Ke7
4 a6        Kd8
5 Rd6        Ra7
6 Rxg6    Rxa6
7 Rg8+    Kc7
8 d8=Q+
again White wins.

238VAR-5
1 d5        Ra8
2 d6        Kf7
3 d7        Ke7
4 a6        Kd8
5 Rd6        Rb8
6 Kf3        Rb4
Purpose: to prevent
White's King from
advancing, and to 
get behind the Rook Pawn.

238VAR-6
1 d5        Ra8
2 d6        Kf7
3 d7        Ke7
4 a6        Kd8
5 Rd6        Rb8
6 Kf3        Rb4
7 a7         Ra4
8 Rxg6        Ra5    
9 a8=Q+    Rxa8
10 Rg8+
winning the Rook.

238VAR-7
1 d5        Ra8
2 d6        Kf7
3 d7        Ke7
4 a6        Kd8
5 Rd6        Rb8
6 Kf3        Rb4
7 a7         Ra4
8 Rxg6        Rxa7
9 Rg8+    Kxd7
10 Rg7+
Rook falling again.

239VAR-1
1 Rh8        a3
2 Rxh7+    Rxh7
3 gxh7
pinned Pawn and
exchanged Rooks
threatens quick win.

239VAR-2
1 Rh8        Rd2+
2 Kf1        Rg2
3 Rxh7+    Kg3
4 g7        Kf3
5 Rh3+    Rg3
6 Rxg3+
 and White wins.

239VAR-3
1 Rh8        Rd2+
2 Kf1        Rd1+
3 Ke2        Rg1
4 Rxh7+    Kg3
5 Rh1!        Rxh1
6 g7
win for White.

239VAR-4
1 Rh8        Rd2+
2 Kf1        Rd1+
3 Ke2        Rg1
4 Rxh7+    Kg3
5 Rh1!        Kg2
6 Rxg1+
again a win for White.

239VAR-5
1 Rh8        Rd2+
2 Kf1        Rd1+
3 Ke2        Rg1
4 Rxh7+    Kg3
5 Rh1!        Rg2+
6 Ke3        Kg4
7 g7        Kf5
uncovering the Rook.

239VAR-6
1 Rh8        Rd2+
2 Kf1        Rd1+
3 Ke2        Rg1
4 Rxh7+    Kg3
5 Rh1!        Rg2+
6 Ke3        Kg4
7 Rh2!        Rg1
8 Kf2        Rg3
9 Rg
forcing an exchange 
of Rooks.

239VAR-7
1 Rh8        Rd2+
2 Kf1        Rd1+
3 Ke2        Rg1
4 Rxh7+    Kg3
5 Rh1!        Rg2+
6 Ke3        Kg4
7 Rh2!        Rg3+
8 Kf2        Rf3+
9 Rg2

239VAR-8
1 Rh8        Rd2+
2 Kf1        Rd1+
3 Ke2        Rg1
4 Rxh7+    Kg3
5 Rh1!        Rg2+
6 Ke3        Kg4
7 Rh2!        Rg3+
8 Kf2        Rf3+
9 Kg1        Re3
10 g7        Re8
11 g8=Q+    Rxg8
12 Rg2+
Black's Rook comes 
off the board.

240VAR-1
1 axb7        Rxb4
2 b8=Q    Rxb8
White must take the 
draw by 3 Rxf2.

240VAR-2
1 b5        Nd6
2 a7        Ra4
3 d6        Nc4+
4Kd3         Ktxb6
with a draw as result.

240VAR-3
1 a7        f1=Q!
2 Rxf1        Rh2+
3 Kc1!        Ra2
4 b5!        Rxa7
5 Rf6+        Kg5
6 Ra6        Rxa6
7 bxa6
White wins, the Knight
being unable to prevent 
the Pawn from reaching
the last square.

240VAR-4
1 a7        f1=Q!
2 Rxf1        Rh2+
3 Kc1!        Ra2
4 b5!        Ra1+
5 Kc2!        Rxf1
6 a8=Q
and White wins.

240VAR-5
1 a7        f1=Q!
2 Rxf1        Rh2+
3 Kc1        Ra2
4 b5!        Ra1+
5 Kc2        Ra2+
6 Kb1        Rxa7
7 Rf6+        Kg7
8 Ra6        Nd6
9 Rxa7+
capturing with check
saves the Pawn and 
wins for White.

240VAR-6
1 a7        f1=Q!
2 Rxf1        Rh2+
3 Kc1        Ra2
4 b5!        Ra1+
5 Kc2        Ra2+
6 Kb1        Rxa7
7 Rf6+        Kg7
8 Ra6        Rxa6
9 bxa6    Nd8
10 a7
and Black is a move
 too late to catch 
the Pawn.

241VAR-1
1 f6!        d2
2 fxg7+
winning a Rook.

241VAR-2
1 f6!        Rxg6
2 Ra8+
and mate next. 

241VAR-3
1 f6!        gxf6
2 g7+
costs a Rook.

241VAR-4
1 f6!        Rg8
2 Rf7     gxf6
3 Rh7#
Black must depend on 
his passed Pawns, since
the alternative here is mate.

241VAR-5
1 f6!        Rg8
2 Rf7     Rd8
3 Rxg7        d2
4 Rh7+    Kg8
5 f7+        Kf8
6 Rh8+        Ke7
7 Rxd8        Kxd8
8 f8=Q+
which is not appetizing with 
quick Pawn elimination.

241VAR-6
1 f6!        Rg8
2 Rf7     d2
3 fxg7+        Rxg7
4 Kxh6        Rg8
5 Rh7#
Not a glimmer 
of hope. 

241VAR-7
1 f6!        Rg8
2 Rf7     d2
3 fxg7+        Rxg7
4 Kxh6        Rxf7
5 gxf7
Black's King must wait 
for the coup-de-grace.

242VAR-1
1 e5        fxe5
2 fxe5        Rc7
3 Rd7+    Rxd7
4 e6+        Ke7
5 exd7        Kxd7
6 Kg6
White captures abandoned 
Pawns and wins.

242VAR-2
1 e5        fxe5
2 fxe5        Ke7
3 e6        Ra6
4 Rd7+    Kf8
5 Kg6!    Rxe6+
6 Kh7        Re4
7 Rxg7    Re6
8 Rg6
Black’s remaining
Pawn falls.

242VAR-3
1 e5        fxe5
2 fxe5        Ke7
3 e6        Ra4
4 g5!        hxg5
5 Rd7+        Kf8
6 Rf7+        Kg8
7 Kg6        g4
8 h6!        Ra6
9 h7+        Kh8
10 Rf8#

242VAR-4
1 e5        fxe5
2 fxe5        Ke7
3 e6        Ra6
4 g5!        hxg5
5 Rd7+        Kf8
6 Rf7+        Kg8
7 Kg6        g4
8 h6!        Ra8
9 hxg7        g3
10 e7        g2
11 Rf8+
 mates next move.  

242VAR-5
1 e5        fxe5
2 fxe5        Ke7
3 e6        Ra6
4 g5!        hxg5
5 Rd7+        Kf8
6 Rf7+        Kg8
7 Kg6        g4
8 h6!        Ra8
9 hxg7        g3
10 e7        Ra6+
11 Rf6        Rxf6+
12 Kxf6
White mates in
 two more moves.

244VAR-1
1 Qd4        Kg2
2 Qg4+    Kh2
3 Qf3        Kg1
4 Qg3+    Kh1!
Black abandons the 
Pawn nonchalantly, 
since taking it allows
a draw by stalemate.

244VAR-2
1 Kf4        Kh1
2 Kg3??    f1=N+
Black wins the Queen. 

244VAR-3
1 Kf4        Kh1
2 Qe2      Kg2
3 Kg4        Kg1
4 Kg3        f1=Q
5 Qh2#

244VAR-4
1 Kf4        Kh1
2 Qe2        Kg2
3 Kg4        Kh1
4 Qf1+    Kh2
5 Qxf2+ 
captures Pawn
and wins.

245VAR-1
1 Qb3        Kd2
2 Qb2        Kd1
3 Kf3!        c1=Q
4 Qe2#

245VAR-2
1 Qb3        Kd2
2 Qb2        Kd1
3 Kf3!        Kd2
4 Kf2        Kd3
5 Ke1
forces Black to give
 up his Pawn.

246VAR-1
1 Qf5        Kb2
2 Qf2        Kb1
3 Kb3              c1=Q
4 Qa2#

246VAR-2
1 Qf5        Kb2
2 Qf2        Kb1
3 Kb3        c1=N+
4 Ka3        Nd3
5 Qd2
and mate next move.

246VAR-3
1 Qf5        Ka1
2 Kb3!        c1=N+
3 Ka3        Nd3
4 Qxd3
stalemate. 

246VAR-4
1 Qf5        Ka1
2 Kb3!        c1=N+
3 Ka3        Nd3
4 Qf1+
and mates next move.

247VAR-1
1 Qb7+        Ka1
2 Qe4        Kb2
3 Qd4+        Kb3
4 Kd3        Ka3
5 Kc4        a1=Q
Black is then mated
on the next move.

247VAR-2
1 Qb7+    Ka1
2 Qe4        Kb2
3 Qd4+    Kb1
4 Qd1+    Kb2
5 Qd2+    Ka1
6 Qc1#

247VAR-3
1 Qb7+    Ka1
2 Qe4        Kb2
3 Qd4+    Kb1
4 Qd1+    Kb2
5 Qd2+    Kb1
6 Kd1!        a1=N
7 Qb4+    Ka2
8 Kd2        Nb3+
9 Kc3        Ka1
10 Qa3# 

247VAR-4
1 Qb7+        Ka1
2 Qe4        Kb2
3 Qd4+        Kb1
4 Qd1+        Kb2
5 Qd2+        Kb1
6 Kd1!        a1=N
7 Qb4+    Ka2
8 Kd2        Nb3+
9 Kc3        Ka1
10 Qxb3 
extinction.

247VAR-5
1 Qb7+        Ka1
2 Qe4        Kb2
3 Qd4+        Kb1
4 Qd1+        Kb2
5 Qd2+        Kb1
6 Kd1!        a1=N
7 Qb4+    Ka2
8 Kd2        Nb3+
9 Kc3        Ka1
10 Kxb3    Kb1
11 Qe1#

249VAR-1
1 Qe4        Kb2
2 Qb4+    Kc2
3 Qa3        Kb1
4 Qb3+    Ka1!
Black has been forced
to block his Pawn, 
but White cannot gain 
a move for his King to 
approach since he must 
release the stalemate.

249VAR-2
1 Kb6!        Kb2
2 Ka5+        Kc1
3 Qh1+    Kb2
4 Qg2+        Kb3
5 Qg7        Ka3
6 Qa1
White winning easily.

250VAR-1
1 Qg8+    Kh1
2 Qg3        a3
3 Qf2        a2
4 Qf1#

250VAR-2
1 Qg8+        Kf2
2 Qh7        Kg3
3 Qd3+        Kg2
4 Qe4+        Kf2
5 Qh1
White wins. 

250VAR-3
1 Qg8+        Kf2
2 Qh7        Kg3
3 Qd3+        Kg2
4 Qe4+        Kg1
5 Qg4+    Kf2
6 Qh3        Kg1
7 Qg3+        Kh1
8 Qf2
and mate next move.

250VAR-4
1 Qg8+        Kf2
2 Qh7        Kg3
3 Qd3+        Kg2
4 Qe4+        Kg3
5 Kc5        a3
6 Kd4        a2
7 Qh1        a1=Q+
8 Qxa1        Kg2
9 Qb2+    Kg3
10 Qb7    Kh3
11 Qh1
is decisive. 

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
	fmt.Println("{")
	for _, match := range matches {
		moves := []string{}
		comment := []string{}
		for _, line := range strings.Split(match[3], "\n") {
			if ok, _ := regexp.Match(`^([0-9]+).+$`, []byte(line)); ok {
				if ok, _ := regexp.Match(`^([0-9]+)\..+$`, []byte(line)); ok {
					moves = append(moves, line)
					continue
				}
				rx := regexp.MustCompile(`^([0-9]+)(.+)$`)
				line = rx.ReplaceAllString(line, "$1.$2")
				moves = append(moves, line)
			} else {
				comment = append(comment, line)
			}
		}
		movesString := strings.Join(moves, "\\n")
		commentString := strings.Join(comment, " ")
		if lastNumber != match[1] {
			fmt.Printf("],\n'%v': [\n", match[1])
			lastNumber = match[1]
		}
		fmt.Printf("{\ninitialBoard: '%v',\nactions: '%v',\ncomment: `%v`\n},\n", initialBoards[lastNumber], movesString, commentString)
	}
	fmt.Println("]\n}")
}

// Use this in tryramda
// mapObjIndexed((v, k) => [...v.slice(0, 1), ...(newContent[k] ? newContent[k] : [])], oldContent)
