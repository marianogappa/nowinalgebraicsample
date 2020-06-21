package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	rx := regexp.MustCompile(`(?s)(\d+)(.*?)###`)
	matches := rx.FindAllStringSubmatch(text, -1)
	fmt.Println("{")

	for _, match := range matches {
		gn := match[1]
		rawContent := match[2]
		lines := []string{}
		for _, line := range strings.Split(rawContent, "\n") {
			if line == "" {
				continue
			}
			if ok, _ := regexp.Match(`^\s*([0-9]+).+$`, []byte(line)); ok {
				rx := regexp.MustCompile(`^\s*([0-9]+)( +)?([^ ]+)( +)?(.+)?$`)
				matches := rx.FindStringSubmatch(line)
				if len(matches) < 4 {
					fmt.Println("This line failed", line)
				}
				lines = append(lines, fmt.Sprintf(`
                    <div class='notationLine'>
                        <i class='notationMoveNumber'>%v</i>
                        <span class='notationWhiteMove'>%v</span>
                        <span class='notationBlackMove'>%v</span>
                    </div>
                    `, matches[1], strings.Replace(matches[3], "...", ". . .", -1), matches[5]))
			} else {
				rx := regexp.MustCompile(`([ (])([\d]+) `)
				line = rx.ReplaceAllString(line, "$1<i>$2</i> ")
				lines = append(lines, fmt.Sprintf("<p>%v</p>", line))
			}
		}
		fmt.Printf("'%v': `%v`,\n", gn, strings.Join(lines, "\n"))
	}
	fmt.Println("\n}")
}

var text = `
###1
PROMOTING the Pawn will win easily for White. With King and Queen against King, the process of forcing mate is elementary.
Can the Pawn advance at once, or will it be caught by Black? Let's try moving the Pawn and see:
1 h4        Kc4
2 h5        Kd5
3 h6        Ke6
4 h7        Kf7
5 h8=Q

White wins
We can thus play out the moves (mentally, of course) or get the answer by a simple count. The latter method tells us that after four moves the Pawn will stand at h7, and Black's pursuing King at his f7—too late to stop the Pawn from Queening
###2
WHITE cannot win by merely advancing the Pawn. A little preparation is needed.
A simple count shows that the Pawn can reach a7 in four moves, but that the pursuing King, rushing back along the diagonal, will head it off at b7.
The play would be: 1 a4, Ke4 2 a5, Kd5 3 a6, Kc6 4 a7, Kb7 and the Pawn is lost. White must try other means:
        1 Kf5!
This prevents Black from moving along the diagonal to stop the Pawn, and threatens to win by 2 a4, 3 a5, etc. 
      1 ...        Ke3
Black seizes a new diagonal and is ready for this continuation: 2 a4, Kd4 3 a5, Kc5 4 a6 Kb6, and he wins the Pawn.
        2 Ke5!
Interferes with that little idea, and once again threatens to win by pushing the Pawn.
      2 ...         Kd3
Now if  3 a4,  Kc4 4 a5, Kb5 and Black draws.
        3 Kd5!
Cuts Black off a third time from pursuing the Pawn.
        3 ...        Kc3
This gives White a last chance to start the Pawn prematurely.  If now    4 a4, Kb4 nails the Pawn down.  
        4 Kc5!
This is decisive! Black's King is shut out completely.  The continuation (regardless of anything Black does, consider for example Black’s move   4 ... Kd3) will be:
    5 a4    Kc3
    6 a5    Kb3
    7 a6    Ka4
    8 a7    Ka5
    9 a8=Q#
Or any other Black moves also result in White’s ultimate victory.
###3
THE key to the win:
Move the King, not the Pawn!
The King must clear a way for the Pawn to come through. He does this by taking the opposition: he moves to a face to face position with Black's King, and forces Black to give way.  White can then seize control of e8, the square on which the Pawn promotes to a Queen.
     Without this preparation, advancing the Pawn does not work. For instance: 1 e6, Kd8 2 Kd7+, Kd8 3Kd6, (to protect the Pawn) and Black draws by stalemate.
     The proper technique:
1 Ke6!
     The Kings face each other, and White having the opposition, forces Black to give way. The King that can compel the other to yield ground either by retreating or stepping aside, is said to have the opposition.
     This is now for White the ideal position in King and Pawn endings. He wins no matter whose move it is
1 ...    Kd8
Or   1 .... Kf8 2 Kd7, Kg7 3 d6+, Kf8  4 e7+, and wins.
2 Kf7
From this square, the King not only controls e8 (the square on which the Pawn becomes a Queen) but is also in position to escort the Pawn on its way up.
    2 ...         Kd7
    3 e6+        Kd8
    4 e7+        Kd7
    5 e8=Q+
White wins
###4
WHITE wins this ending by getting the opposition, and maintaining it.

1 Kf5
Strong, but obviously forced, as his Pawn was in danger.
The Kings face each other with one square between them. White has the opposition, since his opponent must give ground by retreating or stepping to one side.
1 ...        Kf8
2 Kf6!
Best, since it maintains the opposition. White must resist the temptation to move closer to the Pawn. This would be the consequence: 2 Ke6, Ke8 3 d7+, Kd8 4 Kd6, and Black draws by stalemate. Note in this and similar positions that the Pawn must reach the seventh rank without checking in order to win.
2 ...    Ke8
3 Ke6    Kd8
4 d7    Kc7
5 Ke7

White wins
###5
ANOTHER song to the same tune. White uses the force of the opposition to help his Pawn reach the Queening square.
1 Kd5    Kd7
2 Kc5
If Black could meet this by 2 ... Kc7, he would have the opposition. This being impossible, he sets a little trap.
2 ...    Kd8
He hopes that White will play the obvious 3 Kc6 (moving closer to the Pawn) whereupon    3... Kc8 4  b7+, Kb8  5 Kb3 draws.
3  Kd6!
This gives White the opposition and assures the win, since the Pawn will now reach the seventh rank without checking.
###6
WHITE wins this ending by using his King aggressively. The King heads for c6, an ideal King and Pawn position, since it wins with or without the move.
1 Kg6!
The only move to win. Moving    1 c5 instead allows 1... Kf7 2 Kg5, Ke6 and Black will win the Pawn.
After the actual move, White has the opposition and by maintaining it makes the winning process child's play.

    1 ...        Kf8
    2 Kf6        Ke8
    3 Ke6        Kd8
    4 Kd6

Black never gets a chance to approach the Pawn.  

    4 ...        Kc8 
    5 Kc6

An endgame position worth remembering, since White wins no matter whose turn it is to move.

    5 ...        Kb8

If instead 5 ... Kd8 6 Kb7, (seizing control of the Pawn's Queening square) Ke7 7 c5, and the Pawn cannot be stopped. 

    6 Kd7        Kb7
    7 c5        Kb8
    8 c6        Ka7
    9 c7

White wins
###7
WHITE plays to get his King in front of the Pawn. The position is then a win, no matter whose move it is.
1 Kf7!
There is no win after 1 g6+, Kh8   2 g7+, (2 Kf7 stalemates Black) Kg8 3 Kg6, and Black is stalemated. Notice that the Pawn check at the seventh rank does not win unless White controls the Queening square.
1 ...     Kh8
2 Kg6!    Kg8
3 Kh6    Kh8
4 g6    Kg8
5 g7
The winning idea. The Pawn reaches the seventh rank without checking.
5 ...    Kf7
6 Kh7
Seizes control of g8, the Pawn's Queening square.
6 ...     Ke7
7 g8=Q
White wins
###8
WHITE plays to get the opposition. This will place his King in a dominating position, force Black to retreat, and clear a way for the Pawn to come through.
    1 Ke4    Ke6
White's position is ideal. His King is in front of the Pawn with one square between them. In cases such as this, where he does not have the opposition, he can wrest it from Black by gaining a tempo with a Pawn move.
2 e3!
Shows the value of having a square between King and Pawn. The spare move leaves the position of the Kings unchanged—but it is Black's move and he must give way!
    2 ...        Kd6
    3 Kf5        Ke7
If instead 3 ... Kd5, the procedure is 4 e4+, Kd6 5 Kf6!, Ke7 6 e5, Ke8 7 Ke6, and White wins.
4 Ke5
seizes the opposition. Black must retreat or step aside.
4 ...         Kd7
5 Kf6             Ke8
Or 5 ... Kd6 6 e4, Kd7 7 e5, Kd8 (7 ... Ke8 8 Ke6) 8 Kf7 and White wins.
6 Ke6
The simplest, as White has a forced win no matter whose move it is, and no matter whether his Pawn stands on e2, e3, e4 or e5.
6 ...        Kf8
7 e4        Ke8
8 e5         Kf8
9 Kd7        Kf7
10 e6+        Kf8
11 e7+
White wins
###9
TO CLEAR a path for the Pawn, White must maneuver his King to a dominating position in front of the Pawn. He does so interestingly enough by having the King take a walk along the diagonal!
1 Kf2        Kg6
2  Ke3        Kf5
3  Kd4!        Ke6
4  Kf5        Kd7

If instead 4 ... Ke5 5 d4+, Ke6 6 Kf6, (but not 6 d5+, Kd7 7 d6, Kd8 8 Kf6, Kf8 9 d7+, Kd8 and Black draws) Ke7 7 d5, Kd8 8 Kd6 and we have a familiar win.
5 Kd5     Ke7
White could now ruin things by 6 d4, Kd7, and Black has the opposition and draws.

6  Kc6         Ke6
Or 6 ... Kd8 7 Kd6 and White wins.
7 d4         Ke7
8  d5        Kd8

On 8 ... Ke8 9 Kc7 ends it.
9 Kd6
White wins
###10
WITHOUT the aid of the King, who makes a little trip, White's Pawn could never get to the Queening square.
1 Kb3
If instead 1 g4 (starting out unescorted) Kc5 2 g5, Kd5 3 g6, Ke6, and the Pawn is lost.(i.e. 4 g7, Kf7 5 g8=Q,  Kxg8!)
1 ...              Kc5
2 Kc3
Here too, and for the next few moves, advancing the Pawn permits Black to capture it and draw.
2 ...             Kd5
3 Kd3                  Ke5
4 Ke3                        Kf5
5 Kf3                     Kg5
6 Kg3
Sometimes the way to ensure the Pawn's progress is to get right in its way!
6  ...     Kf5
7 Kh4      Kf6
8 Kh5    Kg7
Or 8 ... Kf5 9 g4+, Kf6 10 Kh6, Kf7 11 g5, Kg8 (11 ... Kf8 12 Kh7) 12 Kg6 and White wins.
9 Kg5            Kf7
10 Kh6               Kg8

If 10 ... Kf6 11 g4, Kf7 12 g5 as in the previous note.
11 Kg6 
White wins
###11
IN ORDER to win (and this applies to all endings where the passed Pawn is on the Rook file) White must prevent Black's King from getting in front of the Pawn, since that always ends in a draw. He must also (this may seem strange!) guard against having his own King boxed in!
These are the drawing possibilities:
A] 1 Kb6, Kc8 2 a4, Kb8 3 a5, Ka8 4 a6, Kb8 5 a7+, (nothing else is any better) Ka8 6 Ka6 and Black is stalemated.
B] 1 Kb6, Kc8 2 Ka7, (to prevent 2... Kb8) Kc7 3 a4, Kc8 4 a5, Kc7 5 a6, Kc8 6 Ka8, Kc7 7 a7, Kc8, and White is stalemated.
The way to win is to use the King aggressively—keep Black out of c8 (and any drawing chances) completely.

1 Kb7!
After this move, Black has no play at all.
1     ...    Kd7
2 a4        Kd6
3 a5        Kc5
4 a6
White wins. The Pawn marches gaily up to become a Queen.
###12
THIS looks like an easy win, since the connected Pawns are never in danger.  If the rear Pawn is captured, the other one advances.  . However there remains the danger of stalemating Black.
White solves the problem by remembering that one Pawn ahead is enough to win.

1  Kf 4        Kg7
2  Kf5        Kh8
3 Kg5
If White tries 3 Ke6, then after      3 ... Kg7  4 Kd7, Kh8, his King must not come any closer.

3 ...             Kg7
4  h8=Q+!            Kxh8
5  Kf 6            Kg8
6  g7            Kh7
7 Kf7
White wins
###13
DOUBLED Pawns are often an ad-vantage in the endgame if only for the fact that one of the Pawns may give up his life for the other!
1 g7    Kh7
Careful, now! If White moves 2 Kf7 he allows a draw by stalemate, while 2 g6+, Kg8 lets the win slip forever.
2 g8=Q+!    Kxg8
3 Kg6
To get the opposition.
3 ...    Kh8
Or 3 ... Kf8 4 Kh7 and White wins.
4 Kf7    Kh7
5 g6+    Kh8
6 g7+
White wins
###14
THE isolated Pawns look weak and helpless. They are perfectly safe. An attack on either Pawn is parried by advancing the other!
1 a4!
Now if Black plays 1 ... Kxc3 the reply is 2 a5, and the Pawn cannot be overtaken.
1 ...    Kc5
2 Kg3
But not 2 c4, Kxc4 nor 2 a5, Kb5. It is important not to advance the Pawns prematurely.
2 ...    Kb6
Here too against 2 ... Kc4, the Pawns stay put while the King moves closer by 3 Kf3.
3 c4    Ka5
The reply to 3 ... Kc5 would be 4 a5.
4 c5    Ka6
5 Kf3    Kb7
6 a5    Kc6
7 a6    Kc7
8 Ke4    Kc6
9 Ke5    Kc7
10 Kd5    Kb8
11 c6    Kc7
If 11 ... Ka7 12 c7 wins.
12 a7
White wins
###15
THERE is a lesson in this ending for those who "take first and think it over later."
1 Kf4        Kh7
If instead 1 ... g5+ 2 Kf5, g4 3 Ke6, g3 4 f7, Kg7 (on 4 ... g2, White Queens with check) 5 Ke7 and White wins.
2 Kg5        Kh8
3 Kh6!
Taking the opposition wins; taking the Pawn does not.
3 ...        Kg8
Or 3 ... g5 4 f7 followed by     5 f8=Q#.
4 Kxg6        Kh8
5 Kf7        Kh7
6 Ke8
White wins
###16
1 Kb1!
A CURIOUS way to go after the Pawn! The obvious move 1 Kc3 lets Black escape. For example 1 Kc3, a3! 2 b4 (or 2 bxa3, Ke6, and Black scurrying over to a8 draws easily against a Pawn on the Rook file) Ke5 3 Kb3, Kd5 4 Kxa3, Kc6 5 Ka4, Kb6  6 b5, Kb7 7 Ka5, Ka7 8 b6+, Kb7 9 Kb5, Kb8 10 Kc6, Kc8, and White cannot avoid the draw.
1 ...     a3
On 1 ... Ke5 2 Ka2, Kd5 3 Ka3, Kc5 4 Kxa4, Kb6 5 Kb4, and White having the opposition wins.
2   b3!        Ke5
3   Ka2        Kd5
4   Kxa3        Kc6
5 Ka4!
Certainly not 5 Kb4, Kb6, and all White's earlier strategy has been wasted since Black has the opposition and draws.
5 ...         Kb6
6 Kb4
White wins
###17
BLACK is double-crossed by his own Pawn! Without the Pawn on the board, Black's King simply heads for the magic square a8, after which no power on earth can force a win.
1 a4    Ke4
2 a5    Kd5
3 a6    Kd6
Black's own Pawn stands in his way. His King is unable to move to the vital square c6.
4 a7    Kc7
Too late.
5 a8=Q
White wins
###18
A STRAIGHT line is not the shortest distance between two points.
If White moves across the board to capture the Pawn, he wins the Pawn but not the game.
This is what would happen: 1 Ke7, Kc3 2 Kd7, Kd4 3 Kc7, Kc5 4 Kb7, Kd6 5 Kxa7, Kc7 6 Ka8, Kc8, and Black gets a draw.
Suppose he meanders down the board and then up again?
1 Ke6        Kc3
2 Kd5!
This is the point. Black is kept out of his d4 square.
2 ...        Kb4
3 Kc6        Kc4
4 Kb7        Kb5
5 Kxa7        Kc6
6 Kb8
White wins
###19
WHITE'S King must take a devious route to help his Pawn become a Queen.
1 b5    Ke5
There is no hope in 1... h4 2 b6, h3 3 b7, h2 4 b8=Q+.
2 b6!
If 2 Kc5 (trying to exclude Black's King) Ke6 3 Kc6 (3 b6, Kd7 is a draw) h4 4 b6, h3 5 b7, h2 6 b8=Q, h1=Q and the position is a draw.
2 ...        Kd6
3 Kb5    h4
If 3...Kd7 4 Ka6, Kc8   5 Ka7 wins.
4 Ka6
Of course not 4 b7, Kc7 5 Ka6, Kb8 and White loses.

4 ...        h3
5 b7        Kc7
6 Ka7        h2
7 b8=Q+

White wins
###20
THE race to Queen a Pawn is also a race to see who can check first, and perhaps win the other Queen.
    1 e6    Kf6
Black loses quickly after 1 ... g3 2 e7, g2 3 e8=Q, g1=Q 4 Qg8+, and his Queen comes off the board.
    2 Kd6        g3
    3 e7        g2
If 3 ... Kf7 4 Kd7, and the threat of Queening with check wins for White.
4 e8=Q        g1=Q
5 Qf8+        Kg5
6 Qg8+
White wins the Queen and the game
###21
WHITE'S timing has to be precise in order to Queen his Pawn with a check.
1 Kg5    b5
2 Kf4
To stop Black's Pawn. The alterna-tive 2 h5, b4 leads to both Pawns Queening—and a draw.
2 ...    Ke2
Naturally, if 2 ... b4 3 Ke4 will catch the Pawn.
3 Ke4    Kd2
Here too 3 ... b4 loses the Pawn.
4 Kd4    Kc2
Now if White play 5 h5, Black draws by 5 ... b4 6 h6, b3 7 h7, b2 8 h8=Q, b1=Q.

5 Kc5    Kc3

Ready to meet 6 Kxb5 with 6 ... Kd4  7 h5, Ke5, winning White's Pawn and draw in return. 
6 h5                b4
7 h6                b3
8 h7                b2
9 h8Q+    Kc2
10 Qh2+    Kc1
11 Qf4+    Kc2
12 Qc4+    Kd2
13 Qb3    Kc1
14 Qc3+    Kb1
15 Kc4    Ka2
16 Qa5+    Kb1
17 Kb3    Kc1
18 Qe1#
###22
WHERE both sides Queen their Pawns, the one who first gives check has an advantage that is often de-cisive.
1 f7        h2
2 f8=Q        h1=Q
3 Qf3+        Kg1
Forced, as 3... Kh2 allows 4 Qg3#.
4 Qe3+        Kf1
5 Qc1+        Kg2
6 Qd2+        Kf1
7 Qd1+        Kg2
8 Qe2+        Kg1
9 Kg3!
White wins. Black is curiously helpless to prevent mate.
###23
1 Kc3
THE straightforward 1 g4 leads to 1 ... b5 2 g5, b4  3 g6, b3+  4 Kc3, b2 5 g7, b1=Q 6 g8=Q+, Ka1, and a draw.
1 ...        Ka3
To escort the Pawn through. If instead 1 ... b5 2 Kb4, Kb2 3 g4 wins easily.
2 Kc4        Ka4
3 g4        b5+
4 Kd3!        Ka3

The King must lose a move as 4 ... b4  5 Kc2, Ka3  6 Kb1 is hopeless.
5 g5        b4
6 g6        b3
7 g7        b2
8 Kc2!
Forcing Black into line for a check.
8 ...        Ka2
9 g8=Q+        Ka1
Or 9 ... Ka3 10 Qb3#.
10 Qa8#
###24
1 Kc5!
THE star move to help his own Pawn while making it difficult for Black's to advance.
1 ...    g5
If  1 ...  Kg6 (to cope with White's Pawn) 2 b4, Kf7  3 b5, Ke7 4 Kc6, Kd8  5 Kb7, g5  6 b6, g4  7 Ka7, g3 8 b7, g2  9 b8=Q+ wins for White.
2 b4    g4
3 Kd4    g3
4 Ke3    Kg5
To rescue his Pawn, as chasing after White's is useless.
5 b5
But not 5 Kf3 (to which Black would not meekly oblige with 5... Kh4 6 Kg2 and it's all over) Kf5 6 Kxg3 Ke4 and Black captures the Pawn and draws.
5 ...    Kg4
6 b6    Kh3
If 6 ... g2 7 Kf2, Kh3 8 Kg1 wins.
7 b7            g2
8 Kf2        Kh2
9 b8=Q+
White mates next move with Qg3#.
###25
THE natural move does not always win, even in the most innocent-looking position.
Two ideas suggest themselves to White: capturing Black's Pawn or advancing his own. Neither of them works!
If 1 Kxb7, Kb3 2 Kc6 (2 f4, Kc4 3 f5, Kd5 loses the Pawn) Kc4 3 Kd6, Kd4 4 f4, Ke4 and Black wins the Pawn.
If 1 f4, b5 2 f5, b4, and both sides Queen with a drawn result.
1 Kd6!        Ka3
Or 1 ... b5 2 Kc5, Kb3 3 Kxb5, Kc3 4 Kc5, Kd3 5 Kd5, and the Pawn is safe.
2 Kc5        Ka4
3 f4        b5
4 f5        b4
5 Kc4!
Necessary, as 5 f6, b3 6 f7, b2 7 f8=Q, b1=Q 8 Qa8+, Kb3 9 Qb7+, Kc2 is only a draw.
5 ...        b3
6 Kc3!        Ka3
7 f6        b2
8 f7        b1=Q
9 f8=Q+    Ka4

If 9... Ka2 10 Qa8#.
10 Qa8+        Kb5
11 Qb7+
White wins the Queen and the game
###26
1 Kf5!
STARTING the Pawn instead would be premature: C
1 ...         Ke3
If 1...    c5 2 Ke5, Ke3 3 Kd5, Kd3 4 Kxc5 wins              for White.
2 Ke5        c6
Playing for 3 Kd6, Kd4 4 Kxc6, Kc4 and Black draws. There was no chance in 2 ... Kd3 3 Kd5, c6+ 4 Kc5, Kc3 5 a4, and White wins.
3 a4    Kd3
4 a5
Here too, attacking the Pawn by 4 Kd6 allows 4 ... Kc4 and Black saves himself.
4 ...         c5
5 a6        c4
6 a7        c3
7 a8=Q        c2
White must play carefully now. For example if 8 Qe4+, Kd2 9 Qd4+, Ke2 10 Qc3, Kd1  11 Qd3+, Kc1 12 Kd4, Kb2 13 Qe2  (with the last hope of a win by 13 ... Kb1 14 Kc3, c1=Q+ 15 Kb3, and Black will be mated) Ka1! 14 Qxc2 (nothing else is any better) and Black is stalemated.
8 Qd5+!    Ke2
If Black tries 8 ... Kc3 the following occurs: 9 Qd4+, Kb3 10 Qa1, Kc4 11 Ke4, and the Pawn falls.
Or if 8 ... Ke3 9 Qg2, Kd3 (9 ... c1=Q 10 Qg5+ winning the Queen) 10 Qg5, Kc3 followed by 11 Qc1 is decisive.
9 Qa2!        Kd1
 10 Kd4        c1=Q
 11 Kd3
White mates or wins the Queen.
###27
THE obvious attack on Black's Pawns would lead to the following: 1 Kf7, g5 2 Kg7, Kb3 3 Kxh7, Kc4 4 Kg6, g4 5 Kf5, Kd5 6 Kxg4, Ke6 7 Kg5, Kf7 and Black’s King reaches h8 with an automatic draw against the Rook Pawn.
1 h4!    h5
If  1... h6 2 h5, Kb3   3 Kf7 wins at once, or if 1 ... Kb3 2 Kf7, Kc4 3 Kxg7, h5 4 Kg6, Kd5 5 Kxh5, Ke6 6 Kg6, Ke7 7 Kg7 keeps the Black King at arm's length and wins.
Now comes the point of the position. The natural continuation 2 Kf7 allows 2 ... g5 3 hxg5, h4    4 g6, h3 and both sides Queen with a drawn result
2 Kf8!
This move keeps g8 open so that after 2 ... g5 3 hxg5 White's Pawn will Queen with check.
2 ...    g6
3 Ke7!
Here too if 3 Kf7 or 3 Kg7, the reply 3 ... g5 4 hxg5, h4 leads to a draw. The move actually made keeps the square g8 open for White's Pawn to Queen with check.
3 ...    g5
Or 3 ... Kb3 4 Kf6 and White captures both Pawns and wins.

4 hxg5    h4
5 g6    h3
6 g7    h2
7 g8=Q+    Ka3
8 Qg2
White wins
###28
1 Kf5
WHITE abandons his passed Pawn. Capturing it will keep Black busy on one side of the board, while White gets time to win on the other.
1 ...           Kh6
The Pawn must be removed. If Black tries to defend the Knight Pawn, then after 1 ... Kf7 2 Ke5, Ke7 3 Kd5, Kd7 4 h6 makes him regret it.
2 Ke5                Kxh5
3 Kd5                   Kg6
4 Kc6                  Kf6
5 Kxb6              Ke7
6 Kc7
        White wins
Another way to win is this:
1 h6+                  Kh7
2 Kh5        Kh8
3 Kg6        Kg8
4 h7+        Kh8
5 Kh6        b5
6 a5        b4
7 a6        b3
8 a7        b2
9 a8=Q#
###29
TRYING to win by promoting the Bishop Pawn does not work. For example, if 1 f5, Kf7 2 Ke5, Ke7 3 f6+, Kf7 4 Kf5, Kf8 5 Ke6. Ke8 6 f7+, Kf8 7 Kd6 (too late, but the alternative 7 Kf6 draws by stalemate) Kxf7 8 Kc6, Ke7 9 Kb6, Kd7 10 Kxa5, Kc7 11 Kb5, (or Ka6,) Kb8 , and Black draws against the Rook Pawn.
The way to win is to abandon the Bishop Pawn, keep Black occupied in capturing it, and win on the other side of the board.
1 Kd5        Kf5
2 Kc5        Kxf4
3 Kb5        Ke5
4 Kxa5        Kd6
5Kb6         Kd7 
6 Kb7
Shuts the King out, after which the Pawn has clear sailing.
6 ...        Kd8
7 a5
White wins
###30
TO WIN this, the King forces his way in behind Black's Pawn and attacks from the rear.
1 Kd5!
Of course not 1 Kf5, Kd6 2 Kf6, Kc5 3 Ke6, Kxb5 4 Kd5, Kb6 and White loses both Pawns!
1 ...          Kd8
Black must stay near his Pawn.   If he makes some such move as  1...  Kf7 2 b6 wins instantly.
2 Ke6                Kc8
If 2 ... Ke8 3 b6, cxb6 (or     3 ... Kd8 4 b7 and mate next move) 4 c7 wins.
3 Ke7    Kb8
4 Kd8    Ka7
5 Kxe7    Ka8
6Kd8

White wins
###31
1 Kc5    Kb8
BLACK has two other replies we should look at:
If 1 ... . b6+ 2 Kc6 (not 2 axc6+, Kb7 and Black draws) bxa5 3 Kc7, a4 4 b6+, Ka6 5 b7, a3 6 b8=Q and White wins.
If 1 ... Ka8 2 Kb6, Kb8 3 a6, bxa6    (or    3 ... Ka8 4 axb7+, Kb8 5 Ka6, Kc7 6 Ka7) 4 Kxa6!, Ka8 5 b6, and wins.

2 Kb6        Ka8
3 Kc7        Ka7
4 a6
One advantage in being a Pawn ahead is that you may sacrifice one Pawn for the sake of the other.
4 ...            bxa6
5 b6+            Ka8
6 b7+            Ka7
7 b8=Q#
###32

###33
THE key to the win is for White to give up his passed Pawn, force his way to d6, capture Black's Pawn and then promote the remaining Pawn.
1 Kf6     Kd8!
On 1 ... Ke8 2 Ke6, Kd8 3 d7, Kc7 4 Ke7 wins.
2 d7!    Kxd7
3 Kf7!
White now has the opposition, and will regain the Pawn he sacrificed.

3 ...             Kd8
4 Ke6           Kc7
5 Ke7           Kc8
6 Kd6           Kb7
7 Kd7           Kb8
8 Kxc6         Kc8
9 Kd6           Kd8
10 c6              Kc8
11 c7
###34
AN EXCHANGE of Pawns, after suitable preparation, wins for White. The preparation consists in getting his King behind Black's Pawn, into a dominating position.
1 Kd7
The immediate exchange is premature, as after 1 d5, cxe5 2 Kxe5, Kb6 3 Kc4, Kc6 4 b5+, Kb6, the position is a familiar draw.
1 ...    Kb6
2 Kc 8    Ka6
3 Kc7    Kb5
4 Kb7    Kxb4
5 Kxc6             Kc4
6 d5

White wins
###35
THE method here is to crowd Black into a corner, sacrifice one Pawn and break through with the other.

1 Kd6    Kd8
2 Ke6    Kc8
The alternative loses at once: 2 ... Ke8 3 c6, Kd8 (or 3 ... bxc6 4 b7) 4 cxb7, Ke8 5 b8=Q#.
3 Ke7    Kb8
4 Kd7    Ka8
5 c6
Clearly, not 5 Kc7 stalemating Black.

5 ...            bxc6 
6 Kc7            c5
7 b7+             Ka7
8 b8=Q+            Ka6
9 Qb6#
###36
1 Kd7!
A STRANGE-LOOKING move! From this square though, the King has two wins prepared, depending on which of two possible first moves Black selects.
1 ...         Kb7
Against the alternative 1 ... Ka7 White seizes the opposition, gets behind the Pawn and wins this way: 2 Kc7, Ka6 3 Kb8, Ka5 4 Kb7, Kxa4 5 Kxb6 followed by 6 c5, and all over for Black!
2 a5!    bxa5
If 2 ... Ka6 instead, then 3 axb6, Kxa5 4 Kd6, Kb7 5 c5, Kc8 6 Kc6 and we have a familiar winning position.
3 c5                    a4
4 c6+                       Kb6
5 c7                   a3
6 c8=Q                     a2
7 Qa8

White wins
###37
IT IS obvious that giving up the Queen Pawn will clear the way for the Knight Pawn to come through, but the sacrifice must be prepared properly. The King must assume a most aggressive position—behind Black's Pawn in fact!
1 Kb7
If at once 1 d6, Kc8 2 d7+ (or 2 b6, cxb6 3 Kxb6, Kd7 drawing easily) Kd8 3 Kb7, Kxd7 4 Kb8, Kd6 and Black draws.
1 ...              Kd7
2 Kb8                  Kd8
Or 2... Kd6 3 Kc8, Kxd5 4 Kxc7. 
3 d6!        cxd6
4 b6        d5
5 Ka7        d4
6 b7        d3
7 b8=Q+        Kd7
8 Qb5+ 
White wins Black’s Pawn
and the game.
###38
WHITE can win easily by going after Black's miserable Pawn thus: 1 Kf5, Kf7 (1 ... Kf8 2 Kg6) 2 Kg5, Kb8 3 Kxh5 etc.
Or, he can indulge in a bit of refined cruelty by allowing Black to get a Queen before being mated.
1 f7+        Kf8
2 Kf6        h4
3 g4        h3
4 g5        h2
5 g6        h1=Q
6 g7#
###39
SHOULD White begin with 1 g3 or 1 g4? The Pawn must reach g7 without giving check, that is to say when the opposing King stands at his g8 square.
The right move is then ...
 1 g3!
The aggressive 1 g4 leads to 1 ... Kh8 2 g5, Kg8 3 g6, hxg6 4 hxg6 (if 4 Kxg6, Kh8, and Black draws automatically against the Rook Pawn) Kh8 5 g7+ and a draw.
1 ...    Kh8
2 g4            Kg8
3 g5    Kh8
4 g6    hxg6
Or 4 ... Kg8 5 g7, Kf7 6 Kxh7 and White wins.
5 hxg6    Kg8
6 g7    Kf7
7 Kh7
White wins
###40
BLACK is faced with the task of defending his Pawn, while restraining White's from advancing.
The difficulties would increase, if it were Black's turn to move.
White therefore plays to bring about the position in the diagram, with Black to move. This he does by the process of triangulation.

1 Kf2               Kg6
On 1 ... Ke5 2 g6, Kf6 3 h5, Ke6 4 g7, Kf7 5 h6, and it is clear that Black's King cannot be all over the board at the same time.

2 Ke2    Kf5
3 Ke3    Ke5
4 g6    Kf6
5 h5    Kg7
6 Kxe4
White wins
###41
TO WIN this White must bring about the position in the diagram, with Black to move. He does this by tri-angulation.
1 Kd5    Kc8
2 Kd4!
White is now ready to meet 2 ... Kc7 with 3 Kc5, and follow it up with 4 Kb6.
2 ...    Kd8
3 Kc4
Still waiting for 3 ... Kc7.
3 ...        Kc8
4 Kd5        Kc7
The alternative is 4... Kd8 5 Kd6, Kc8 6 c7, Kb7 7 Kd7, Ka7 8 Kc6 (but not 8 c8=Q stalemate) Ka8 9 c8=Q+ and mate next.
5 Kc5
Now we have the position in the diagram—but it is Black to play. The rest is easy.
5 ...        Kc8
6 Kb6        Kb8
7 Kxa6        Kc7
8 Kb5        Kc8
9 Kb6        Kb8
/0 c7+        Kc8
11 a6        Kd7
12 Kb7
White wins
###42
THE problem for White is to force an exchange of Pawns. This would transform his two Pawns against one, to one Pawn against none.
1 c5                   Kc8
If 1 ... b6 2 c6, Kc8 3 Kxb6 and it's all over.
2 Kb6    Kb8
3 c6                 Ka8
Hoping for 4 b7 stalemate. If instead 3 ... Kc8 4 c7, Kd7 5 Kxb7 wins, or if 3 ... bxc6 4 Kxc6, Kc8 5 Kb6, Kb8 6 b5, Kc8 7 Ka7 wins.
White can now win by 4 cxb7+, Kb8 5 Ka6, or by the following:
4 Kc7                bxc6
5 Kxc6    Ka7
6 b5                Kb8
7 Kb6

Definitely not 7 b6, Kc8 and Black escapes with a draw.

7 ...        Ka8
8 Kc7    Ka7
9 b6+
White wins
###43
WHITE'S strategy is simple: he forces an exchange of Pawns, (at the right time of course) moves his King to a dominating position, and wins.
1 g3
Not at once 1 f3+ as after 1 ... gxf3 2 gxf3+, Kf4, Black has a draw.
1 ...        Kd4
On 1 ... Ke5, White goes after the Knight Pawn as follows: 2 Ke3, Kf5 3 Kd4, Kf6 4 Ke4, Kg5 5 Ke5, Kg6 6 Kf4, Kh5 7 Kf5, Kh6 8 Kxg4 and wins.
2 f3        gxf3+
If 2 ... Ke5, White demonstrates the Queening of a doubled Pawn by this: 3 fxg4, Ke4 4 Kf2, K-e5 5 Ke3, Ke6 6 Kf4, Kf6 7 g5+, Kg6 8 Kg4, Kg7 9 Kf5, Kf7 10 g6+, Kg7 11 Kg5, Kg8 .12 Kf6, Kf8 13 g7+, Kg8 14 g4, Kh7 15 Kf7 and wins.
3 Kxf3        Ke5
4 Kg4!
The King must be aggressive. There is no win if the Pawn moves prematurely.
4 ...        Kf6
5 Kh5        Kf5
Or 5 ... Kg7 6 Kg5, Kf7 7 Kh6 etc.
6 g4+        Kf6
7 Kh6        Kf7
8 g5        Kg8
9 Kg6
White wins
###44
WHITE wins this by beautiful timing of moves.
1 Kb3
With the very first move White can go wrong. For instance, if 1 Kb4, c5+ 2 dxc5+, Kc6 3 Kb3, Kxc5 and the position is a draw. Now if 1 ... c5, White can reply 2 d5 avoiding the exchange of Pawns.
1 ...        Kc7
2  Kc3        Kd6
3 Kd3        Kd7
4 Ke4     Ke6
5  c5        Kf6
If 5 ... Kf7 6 Kf5, Ke7 7 Ke5, Kd7 8 Kf6, Kd8 9 Ke6, Kc7 10 Ke7, Kc8 11 Kd6, Kb7 12 Kd7 and Black's Pawn falls.
6 d5!        Ke7
On 6 ... cxd5 7 Kxd5, Ke7 8 Kc6, Kd8 9 Kb7 assures the Pawn's future.
7 d6+
After 7 dxc6, Kd8 8 Ke5, Kc7
9 Kd5, Kc8 10 Kd6, Kd8 11 c7+, 
Kc8 and Black has a draw.
7 ...          Kd7
8 Ke5          Kd8
9 d7!
The point of the previous maneuvering.
9 ...         Kxd7
Avoiding the capture loses faster: 9 ... Ke7 10 d8=Q+, exd8 11 Kd6, Kc8 12 Kxc5, and we have a standard winning position.
10 Kf6
The King circles about to get behind the Pawn.
10 ...        Kd8
11 Ke6        Kc7
12 Ke7        Kc8
13 Kd6        Kb7
14 Kd7        Kb8
15 Kxc6        Kc8
16 Kd6        Kd8
17 c6        Kc8
18 c7
Advances to the seventh without check, rendering the win certain.
18 . .     .Kb7
19 Kd7
White wins
###45
1 Kf1        Kd4
2 Kf2        Kc5
3 e4!
THIS diverts Black from pursuing the Rook Pawn. If now 3 ... Kb6 4 e5, Kxa6 5 e6 and White's Pawn cannot be caught.
3 ...        Kd4
4 Kf3        Ke5
5 Ke3        Ke6
6 Kd4        Kd6
7 e5+        Ke6
8 Ke4        Ke7
9 Kd5        Kd7
10 e6+        Ke7
11 Ke5
The tempting 11 Kc6 (going after the Rook Pawn) is premature as Black takes the King Pawn without loss of time (one tempo). The continuation would be 11 ... Kxe6 12 Kb7, Kd7 13 Kxa7, Kc7 14 Ka8, Kc8 and White cannot extricate his King, hence a draw.
11 . .        Ke8
12 Kd6        Kd8
13 Kc6
Now it will take Black two moves to capture the Pawn.
13 ...        Ke7
Or quick suicide by 13 ... Kc8
14 e7 and mate next move.
14 Kb7        Kxe6
15 Kxa7    Kc7
16 Kb7
White wins
###46

###47
IT IS clear that White must capture the Knight Pawn in order to win. But how does he do so? The obvious 1 Kg4, Kh2 2 Kg5, Kg3 3 Kg6, Kxf4 leads to a draw.
1 f6!        gxf6
2 f5
Fixing the target. Now the King circles about and picks off the helpless victim.
2 ...        Kh2
3 Kg4        Kg2
4 Kh5        Kf3
5 Kg6        Kg4
6 Kxf6        Kh5
7 Kg7
White wins
###48
1 Kb5        Kc3 
2 Kc5        Kd3 
3 Kd5        Ke3
4 Ke5        Kf3
5 Kf5        Kg3

AND now, White does not play 6 Kg6 which allows 6 ... Kxh4 and a draw, but with ...

6 h6!         gxh6 
7 h5         Kh4 
8 Kg6         Kg4 
9 Kxh6         Kf5
10 Kg7
White wins
###49
THE natural line of play would seem to be to go after the Knight Pawn by way of f5 and g6. Here is what would happen: 1 Kf5, Kf3 2 Kg6, Kg3 3 h5, Kh4 4 Kxg7 (there is nothing better) Kxh5 with a draw as result.
The winning idea is for White's King to reach g6 in a roundabout way—moving to e6, f7, and then g6!
1 Ke6!        Kf4
2 Kf7        Kg3
The desperate 2 ... g5 leads to an easy White win after 3 h5, Kg3 4 Kg6.
3 h5        Kh4
4 Kg6!
Now we have the position after Black's third move in the first paragraph—but this time with Black to play.
4...        Kxh3
5 Kxg7
White wins
###50
TO WIN this White must sacrifice—and at once, or he never gets another chance to do so.
1 e6!
Against any other move Black's reply of 1 ... Ke7 assures him of a draw.

1 ...      fxe6
Black might refuse to capture, with this result: 1 ... Ke7 2 exf7, Kxf7 3 Kd5, Ke7 4 Ke5 and White with the opposition wins. 
Or if 1 ... f6 2 Kc5, Ke7 3 Kd5, Ke8 4 Kd6, Kd8 5 e7+, Ke8 6 Ke6, f5 7 Kxf5, Kxe7 8 Ke5, and again the force of the opposition wins.

2 e5!
Fixes the Pawn. Now the King plays to get behind the Pawn.
2 ...             Ke7
3 Kc5        Kd7
4 Kb6        Kd8
5 Kc6        Ke7
6 Kc7            Ke8
7 Kd6        Kf7
8 Kd7        Kf8
9  Kxe6            Ke8
10 Kf6        Kf8
11 e6        Ke8
12 e7
White wins
##51
THE doubled Pawns find a unique means of keeping the King at bay.
1 a5        Kc5
2 a4        Kd6
Thwarted, the King tries coming around in front of the Pawns.
3 Kd8    c5
4 a6        Kc6
5 a5
Once again the Pawns set up a barricade.
5 ...              c4
6 Kc8              c3
7 a7              c2
8 a8=Q+          Kd6
9 Qh1

White wins
###52
THE doubled Pawns hold their own against the King until help arrives.
1 Kf 1        Kc3
2 Ke2        Kd4
3 g4        Ke4
4 g3        Kd4
5 g5        Ke5
Black must stop the passed Pawn even at the cost of deserting his own Pawn.
6  g4
Prevents the King from coming closer.
6 ...        Ke6
7 Kxe3        Kf7
8 Ke4
Not at once 8 Kf4 as the reply 8 ... Kg6 costs a Pawn and gives Black an easy draw.
    8 ...        Kg6
Or 8 ... Kg6 9 Kf4 followed by 10 Kf5.
9 Kf4        Kf7
    10 Kf5        Kg7
    11 g6        Kg8
If 11 ... Kh6 12 g7 (certainly not 12 Kf6 stalemating Black) Kxg7 13 Kg5, and White having the opposition wins.
12 Kf6        Kf8
    13 g7+        Kg8
14 g5        Kh7
15 g8=Q+
Here too, the hasty 15 Kf7 stalemating Black, deprives White of the win.
15 ...        Kxg8
16 Kg6
White wins
###53
1 Kg3!
THIS apparent loss of time is the only way to a win. If instead 1 Kg5, Kf3 2 h4, Kg2 3 h5, Kh3! 4 Kg6, Kg4! (not 4 ... Kh4 5 h3, Kxh3 6 Kxg7 and White wins) 5 h3+, Kh4 6 Kxg7, Kxh5(h5) and Black draws. Or if 1 Kg4, Kf2 2 h4, Kg2 3 h3, g6 4 Kg5, Kxh3, and the position is drawn.
1 ...         Ke3
If 1 ... Kf1 2 h4, g6 (or 2 ... Ke2 3 h5, Ke3 4 h4, Ke4 5 Kg4, Ke3  6 Kf5, Kf3  7 h6, gxh6  8 h5 and White wins.) 3 Kf4!, Kg2 4 h5!, gxh5 5 h4 and White wins.
2 h4        Ke4
3 Kg4    Ke5
4 Kg5    Ke4
5 h5        Kf3
6 Kf5!
Not at once 6 Kg6 when 6 ... Kg4 7 h3+, Kh4 ends in a draw.
White's actual move forces Black's King to retreat.
6 ...        Kg2
7 Kg6
White wins the Pawn and the game.
###54
WHITE has a great advantage in his outside passed Pawn. Its threat to march up the board keeps Black's King occupied, since he must eventually chase after it. Meanwhile White has time to pick off the abandoned Black Pawns.
1 a5                  Kc6
2 a6                 Kb6
3 a7                 Kxa7
4 Kxc5     Kb7
5 Kd6       Kc8
6 Ke6      Kd8
7 Kxf5             Ke7
8 Kg6     Ke6
If 8 ... Kf8 9 Kf6 is a win with or without the move.

9 f5 +     Ke7
10 Kg7
###55
IT IS important to Queen a Pawn on the right square, if there is a choice. Promoting on the wrong square might let the opponent get away with a draw when he should have lost.
1 a4    h5
The alternative 1 . . Kf7, trying to head off White's Pawns, is too slow, viz: 2 a5, Ke7 3 b6, axb6 (3 ... Kd7 4 bxa7) 4 a6!, and White wins.
2 a5    h4
3 b6                    axb6
4 a6!

If  4 axb6, h3 5 b7, h2 6 b8=Q, h8=Q and the position is a draw.
4 ...    h3
5 a7    h2
6 a8=Q
White wins
###56
WHITE'S strategy is based on the fact that his Pawns are safe from capture, while Black's will both be lost if either of them makes a move.
1 Kf4    Kb6
The Pawns must not move. If for example 1 ... h5 2 Kg5, Kb6 (or 2 ... d5 3 Kxh5, d4 4 Kg4, d3 5 Kf3 and White overtakes the Pawn) 3 Kxh5, Kc7 4 Kg5, Kb6 5 Kf5, Kc7 6 Ke6, and the second Pawn goes.
2 Kf5    Kc7
3 Kf6    Kb6
Here too if 3... h5 4 Kg5 or if 3... d5 4 Ke5 and White will be in time to catch the second Pawn.
4 Ke6        Kc7
If 4 ... h5 5 Kxd6, h4 6 c7, h3 7 c8=Q, h2 8 Qa6#.
5 Kd5        h5
6 b6+        Kxb6
7 Kxd6        h4
8 c7        Kb7
9 Kd7
White wins
###57
1 Kf7!        h5
ON 1 ... Kh8 instead, 2 Kg6 wins both of Black's Pawns.
2 h4!
The key move to break through. White cannot win with 2 Kf6 as 2 ... h5xg4 3 h3xg4, Kh8 4 Kxg5, Kg7 gives Black the opposition and an easy draw.
2 ...                Kh6
What else is there? if 2 ... h5xg4 3 h4xg5, g3 4 g6+, Kh6 5 g7, g2 6 g8=Q wins, or if 2 ... g5xh4 3 g5, h3 4 g6+, Kh6 5 g7, h2 6 g8=Q, h1=Q 7 Qg6#.
 3 Kf6!        hxg4
Or 3 ... g5xh4 4 g5+, Kh7
5 Kf7 and White wins.
4 hxg5+            Kh5
5 g6            g3
6 g7            Kg4
7 g8=Q+            Kf3
8 Kg5            g2
9 Kh4            Kf2
10 Qg3+            Kf1
11 Qf3+            Kg1
12 Kh3            Kh1
13 Qxg2#
###58
1 Kf4!
UNEXPECTED, but the obvious idea of immediately going after the Rook Pawn leads only to a draw: 1 Kf3, Kg8 2 Ke3, Kf8 3 Kd3, Ke7 4 Kc3, Kf6 5 Kb3, Kxg6 6 Kxa3, Kf5 7 Kb4, g5 8 a4, and both sides will Queen at the same time.
1 ...        Kg8
2 Ke5        Kf8
3 Kd6!    Ke8
4 Ke6        Kf8
On 4 ... Kd8 5 Kf7 is fatal.
5 Kd7        Kg8
6 Ke7        Kh8
Now we see the reason for the tour White's King made. It will take only four moves now to capture Black's Rook Pawn instead of the six it would have taken from the original position. Black's King meanwhile has been unable to make any progress, having been forced back to the corner.
7 Kd6        Kg8
8 Kc5        Kf8
9 Kb4        Ke7
 10 Kxa3
Black is in trouble. If he tries to get a passed Pawn, this is what happens: 10 ... Kf6 11 Kb4, Kxg6 12 a4, Kf5 13 a5, g5 14 a6, g4 15 a7, g3 16 a8=Q, Kf4 17 Qg2 and White wins.
    10 ...        Kd6
So he tries another defense.
    11 Kb4        Kc6
    12 a4        Kb6
    13 Kc4        Kc6
    14 a5        Kc7
    15 Kd5
White wins, since Black cannot stop the Rook Pawn and defend his own Pawn at the same time.
###59

###60
1 Kc4!    Kxf4
NECESSARY, as 1 ... a4 2 d5, a3 3 Kb3 is hopeless for Black.
2 d5        Ke5
3 Kc5        f4
Queening the Rook Pawn instead would result in this: 3... a4 4 d6, a3 5 d7, a2 6 d8=Q, a1=Q 7 Qh8+, and White wins the Queen on the diagonal.
4 d6        f3
Or 4 ... Ke6 5 Kc6, f3 6 d7, f2 7 d8=Q, f1=Q 8 Qe8+, Kf6 9 Qf8+, and White wins the Queen.
5 d7        f2
6 d8=Q    f1=Q
7 Qc8+    Kf5
8 Qf8+
White wins the Queen on the file.
###61
A CLEVER sacrifice transforms a drawish-looking position into a pretty win.
1 Kd4!
Not at once 1 c5, as after 1 ... Ke6 2 Kd4, a4 3 Kc4, a3 4 Kb3, Kd5, and White loses instead of winning.
1 ...        Ke6
On 1 ... a4 instead, 2 Kc3, Ke6 3 Kb4 wins for White.
2 Kc5        Ke5
3 Kb5        Kd4
Expecting the continuation 4 c5, a4 5 c6, a3 with a comfortable draw. Instead of this, Black gets a rude shock.
4 g5!        fxg5
5 c5        a4
Black can Queen the other Pawn, with this result: 5 ... g4 6 c6, g3 7 c7, g2 8 c8=Q, g1=Q  9 Qc5, and White wins the Queen.
6 c6    a3
7 c7    a2
8 c8=Q    a1=Q
9 Qh8+
White wins the Queen and the game.
###62
AN APPARENTLY simple position, but how does White proceed? If 1 Kd8, (to attack the Pawns) Kc6 2 Ke7, f5! 3 gxf5, g4, and both sides will Queen their Pawns with a drawn result.
1 a5!        Kc6
2 Kb8!
Threatens to advance the Pawn and Queen with check.
2 ...        Kb5
3 Kb7        Kxa5
Or 3 ... f5 4 a6, fxg4 5 a7, g3 6 a8=Q, g2 7 Qa1 and White wins.
4 Kc6        f5
The best chance. If instead 4 ... Ka6 5 Kd5, f5 6 gxf5, g4 7 Ke4 and Black's Pawn is lost. Or if 4 ... Kb4 5 Kd5, f5 6 gxf5, g4 7 f6, and the Pawn reaches f8 with check winning for White.
5 gxf5        g4
6 f6            g3
7 f7            g2
8 f8=Q        g1=Q
9 Qa3#!
###63
BLACK'S own King Rook Pawn loses the game for him. If White could only be induced to capture it, the position would be a draw!
1 a3!
The attractive 1 Ke5 instead leads to 1 ... h5 2 c5, Kb5! (but not 2 ... h4 3 c6, h3 4 c7, h2 5 c8=Q, h1=Q 6 Qc4+, Ka3 7 Qb3#) 3 Kd6, h4 4 c6, h3 5 c7, h2 6 c8=Q, h1=Q 7 Qc5+, Ka6! and a draw.
1 ...        h5
2 Kg3        h4+
3 Kh3!
It is important not to remove the Pawn.
3 ...        Kxa3
4 c5        a4
5 c6        Kb2
6 c7        a3
7 c8=Q        a2
8 Qb7+        Kc2
9 Qc6+        Kb2
10 Qb5+        Kc2
11 Qa4+        Kb2
12 Qb4+        Kc2
13 Qa3        Kb1
14 Qb3+        Ka1
The threat of stalemate would allow Black to draw—if he didn't have an extra Pawn!
15 Kg4!        h3
16 Qc2        h2
17-Qc1#
###64
IN A position where the moves seem to be forced, it is easy to play mechanically—and let a sure win slip.
1 b6        b3
2 b7        b2
3 b8=R!
The unthinking promotion to a Queen instead allows 3... b1=Q 4 Qxb1 and with locked Pawn and King -- stalemate.
3 ...        Ka2
4 Ka4        b1=Q
5 Rxb1        Kxb1
6 Kb3
White wins. He will capture Black's Pawn and proceed to Queen his own, while his opponent watches helplessly.
###65
1 Ke4
SEIZES the opposition and forces Black to give way. He has two plausible defenses:
A] 1 ...        Kd6 
2 Kd4
The attractive alternative 2 Kf5 is premature. Starting with that move it takes White seven more moves to capture Black's Rook Pawn and then Queen his own. In the meantime, Black captures on the Queen side, Queens his Pawn, and draws.
2 ...        Ke6
On 2 ... Kc6 White plays 3 Ke5 and after the reply 3 ... Kb6 can move 4 Kd5, force the capture of the Knight Pawn and win, or he can switch to the King side by 4 Kf5, capture the Rook Pawn and win.
3 Kc5        Kf5
4 Kxb5        Kg4
5 Kc4        Kxh4
6 b5        Kg3
7 b6        h4
8 b7        h3
9 b8= Q+        Kg2
10 Qg8+        Kf2
11 Qd5
Threatens to continue with 12 Qh8.
11 . .            Kg1
12 Qg5+        Kf2
13 Qh4+        Kg2
14 Qg4+        Kh2
15 Kd4        Kh1
16 Qxh3+ and wins
B] 1 ...        Kf6
2 Kf4!
White does not attack by 2 Kd5.
It would take eight moves to capture the Knight Pawn and Queen his own. In the same number of moves Black would capture the Rook Pawn, Queen his passed Pawn, and draw.
2 ...        Kg6
3 Ke5!
White wins. He has the pleasant choice of winning on either side of the board, since he can force the capture of either of Black's Pawns.
###66
A BEAUTIFUL example of distant opposition. The Kings are separated by the length of the board, but White's, believe-it-or-not, exerts irresistible pressure on Black's (opposition)!
1 Ke7!
White seizes the opposition, and for the rest of the play leaves Black with only one reasonable move in reply.
1 ...    Ke2
If Black leaves the King file, White wins by attacking on the same file. For instance, if 1 ... Kd2 2 Kd6, followed by 3 Kc5. Or if 1 ... Kf2 2 Kf6, followed by 3 Kg5.
2 Ke6!        Ke3
If 2 ... Kd3 3 Kd5 and 4 Kc5 wins, or if 2 ... Kf3 3 Kf5 and 4 Kg5 wins.
3 Ke5!        Ke2
4 Ke4!        Ke1
Here too 4 ... Kd2 is met by 5 Kd4 and 6 Kc5, while 4 ... Kf2 succumbs to 5 Kf4 followed by 6 Kg5
5 Ke3!        Kf1
Unfortunately, Black must leave the file. There is no square left on it for retreat.
6 Kf4        Ke2
7 Kg5        Kd3
8 Kxh5        Kc4
9 Kg4        Kxb4
10 h5        Ka3
11 h6        b4
12 h7        b3
13 h8=Q        b2
14 Qc3+        Ka2
15 Qc2        Ka1
16 Qa4+        Kb1
The idea in this sort of position is to force Black to block his Pawn. Each time he does so, White gains time to bring his King closer.
17 Kf3        Kc1
18 Qc4+        Kd2
19 Qb3        Kc1
20 Qc3+        Kb1
21 Ke3        Ka2
22 Qc2        Ka1
23 Qa4+        Kb1
24 Kd3        Kc1
25 Qc2#
###67
1 a4        Kd6
BLACK had no time to start his passed Pawn as White's Rook Pawn threatened to move up the board and Queen with check.
2 Kb6
Otherwise Black heads off the passed Pawn with 2 ... Kc7.
2 ...    Kd7
If he can get to c8, White's Rook Pawn will be harmless.
3 Kb7        h5
4 a5        h4
5 a6        h3
6 a7        h2
7 a8=Q        h1=Q
This might be a draw, except that in Queening first White has the advantage of giving a few checks.
8 Qc8+        Kd6
9 Qc6+        Ke5
10 f4+
White wins the Queen by discovered attack.
###68
1 a5    Kc5
2 Kb2!    f5
IF BLACK attacks the Rook Pawn, the play goes like this: 2 ... Kb5 3 Kc3, Kxa5 4 Kd4, Kb5 5 Ke4, Kc6 6 Kf5, Kd7 7 Kxf6, Ke8 8 Kxg5, Kf7 9 Kh6, Kf6 (or 9... Kg8 10 Kg6) 10 g5+, Kf7 11 Kh7 and White wins.
3 gxf5        g4
Hoping to promote at the same time as White, and draw.
4 f6!
Threatens to reach f8 and Queen with check.
4...        Kd6
To stop the Pawn if it takes an- other step.
5 a6        g3
Black has no time to play 5... Ke6 as 6 a7, g3 7 a8=Q wins.
6 f7
Renewing the threat of Queening with check.
6 ...         Ke7
7 a7        g2
Will both sides Queen their Pawns simultaneously?
8 f8=Q+    Kxf8
9 a8=Q+    Ke7
10 Qxg2
White wins. It is just a coincidence that the composer of this ending is named Chekhover.
###69
1 f4
TO THIS Black cannot reply 1 ...  Pd5 as after 2 f5, d4 3 f6, d3 4 f7, d2 White's Pawn promotes to a Queen with check, and wins.
1 ...        Kb4
2 h4
Now if Black moves 2 ... Kc5 3 h5, and the Rook Pawn cannot be caught.
2 ...        d5
The best defense: if White continues by 3 h5, then 3 ... d4 and Black will Queen with check.
3 f5!
To meet 3 ... d4 with 4 f6, threatening to Queen with check.
3 ...        Kc5
4 h5
Stops the King from coming closer, the reply to 4... Kd6 being 5 h6 followed by 6 h7 and 7 h8=Q.
4 ...        d4
Black in turn is ready to refute 5 h6 with 5 ... d3.
5 f6
Renews the threat of Queening with check.
5 ...    Kd6
Stops the Bishop Pawn but now comes... .
6 h6
Preventing 6... Ke6 as then 7 h7 wins.
6 ...        d3
Now White cannot play 7 h7, so he threatens again to Queen with check.
7 f7!        Ke7
8 h7        d2
9 f8=Q+        Kxf8
10 h8=Q+        Ke7
11 Qd4
White wins. A beautiful composition, with Black's alternation of King and Pawn moves the only possible defense to White's seesawing Pawns.
###70
THE outside passed Pawn keeps Black busy on one side of the board to the neglect of the other. White's King takes the opportunity to come down on the unprotected Pawns like a wolf upon the fold.
1 a5                    Kb5
Sooner or later the Pawn must be removed. If instead 1 ... g5 2 hxg5, fxg5 3 Ke5, and Black must still play to capture the passed Pawn.
2 Kd5        Kxa5
3 Ke6        f5
4 gxf5        gxf5
5 Kxf5        Kb6
6 Kg6        Kce7
7 Kxh6        Kd7
8 h5        Ke7
9 Kg7
White wins
###71
WHITE, a Pawn ahead, plays to exchange Pawns on the King side. This would leave a passed Pawn which Black would have to keep under surveillance. White meanwhile could switch his attack to the other side of the board and win the game there.
1 g4        a5
2 a4        Kf6
3 h4        Ke6
On 3 ... Kg6 4 Kd5, Kf6 5 Kc5, Ke5 6 Kb5, Kf4 7 Kxa5, Kxg4 8 Kb6, and White will Queen his Pawn.
4 g5        Kf7
5 Kf5        Kg7
6 h5        Kf7
If 6 ... h6 7 g6, f8 8 Ke6, and the King goes over to the Queen side, since his King side Pawns are safe from capture.
7 Ke5        Ke7
8 g6        hxg6
9 hxg6        Ke8
On 9 ... Kf8, White can win it on the King side if he wishes, by 10 Kf6, since the Pawn advances to the seventh rank without checking.
10 Kd5        Kf8
Black's King, unfortunately for him, cannot be in two places at the same time.
11 Ke5 
White wins
###72
WHITE has three Pawns to two. An exchange of Pawns, leaving him with two Pawns to one, would greatly increase his advantage. By means of a temporary sacrifice, White brings about an exchange.
1 g4+!                  hxg4+
There is no hope in 1 ... Kg6 2 g5, and White has a passed Pawn on the King side.
2 Kg3                 Kg6
If 2 ... Ke5 3 Kxg4, Kd4 4 h5, Kc4 5 h6, and the Pawn Queens long before Black can promote his Pawn.
3 Kxg4                  Kh6
4 Kf5
White can go gaily about his business, since Black must lose time disposing of the passed Pawn.
4 ...             Kh5
5 Ke5                Kxh4
6 Kd5                Kg5
7 Kc5                Kf6
8 Kxb5            Ke7
9 Kc6                Kd8
10 Kb7               Kd7
11 b5
White wins
###73
A CLEVER break-through by White's Pawns!
1 c6!        Kb6
If 1 ... dxc6 instead, then 2 d6 forces 2... exd6, after which White plays 3 f5 and the Pawn marches merrily up the board.
2 d6!        exd6
Other moves lose instantly:
A]
2
. . Kxc6
3 dxe7
B]
2 .
. . dxc6
3 dxe7
C]
2 .
. . e6
3 cxd7


3 f5        Kc7
He must try to stop the passed Pawn.
4 f6        Kd8
5 c7+        Kxc7
The unhappy King cannot be in two places at once.
6 f7 easily Queens next.
White wins
###74
1 g6!
THE only way to break through! White threatens 2 gxh7, winning on the spot.
1 ...        h6
The alternative 1 ... hxg6 leads to this: 2 hxg6, Kf8 3 Kd6 (diagonal opposition) Ke8 4 Ke6, Kf8 (if 4... Kd8 5 Kf7 ends the struggle) 5 Kd7, Kg8 6 Ke7, Kh8 7 f6, gxf6 8 Kf7, f5 9 g7+ and White mates in two.
2 Kd5
Not at once 2 f6+, gxf6+ 3 Kf5, Kf8 4 Kxf6, Kg8 5 g7, Kh7, and Black has a draw.
2 ...        Kf6
On 2 ... Kd7 3 f6 is decisive.
3 Ke4        Kg5
4 Ke5        Kxh5
5 Ke6        Kg5
6 f6!
Now!
6 ...         gxf6
Or 6... Kxg6 7 f7 followed by 8 f8=Q.
7 g7 escapes to Queen.
White wins
###75
1 Kb4                    Kd4
2 a4                       Kd5
3 a5                   bxa5+
IF 3 ... Kd6 instead, 4 axb6 wins at once.
4 Ka4!
Unexpected, but the only way to win. After 4 Kxa5, Black forces a draw by 4 ... Kc5 5 Ka4, Kb6.
4 ...     Kc5
5 Kxa5
Now we have the position in the previous note (after 4 ... Kc5) but with Black to move.
5 ...        Kd6
6 b6        axb6+
No better is 6... Kc6 when 7 bxa7 is conclusive.
7 Kxb6
White wins
###76
BOTH sides have Pawns which are immune from capture. White's are further advanced, so he can venture on a combination.
1 Ke2        Kb7
2 Kd3        Ka8
3 Kc4        Kb7
Black must not push on by 3 ... f3 as the reply 4 Kd3 winning the Pawn is the penalty.
4 Kc5        f3
5 Kd6        f2
6 a8=Q+        Kxa8
7 Kc7        f1=Q
8 b7+        Ka7
9 b8=Q+        Ka6
10 Qb6#
###77
1 Ke4        Kg4
OBVIOUSLY if 1 ... Kxh2 2 g4, and the passed Pawn can never be caught.
2 h4
And now the Knight Pawn may not be captured.
2...        Kh5
3 Kf4        Kh6
4 g4        Kg6
5 h5+        Kh6
6 Ke4        Kg5
7 Kf3        Kh6
8 Kf4        Kh7
9 g5        Kg7
10 g6        Kh6
11 Kg4        Kg7
Naturally, if 11 . . d3 12 Kf3 will win the Pawn.
12 Kg5!        d3
13 h6+        Kg8
14 Kf6         d2
 15 h7+         Kh8
16 Kf7         d1=Q
17 g7+        Kxh7
18 g8=Q+         Kh6
19 Qg6#
###78
1 a3!        f5
BLACK goes straight for a Queen. If he attacks the Rook Pawns instead, the play would go as follows: 1 ... Ka5 2 Kb7, Kxa4 (or 2 ... f5, changing plans in mid-stream, 3 gxf5, g4 4 f6, g3 5 f7, g2 6 f8=Q, g1=Q 7 Qb4#) 3 Kc6, Kxa3 4 Kd5, Kb4 5 Ke5, Kc5 6 Kf6, Kd6 7 Kxf7, Ke5 8 Kg6, Kf4 9 Kh5, Ke5 10 Kxg5, Ke6 11 Kg6, Ke7 12 g5, Kf8 13 Kh7 and White wins.
2 gxf5        g4
3 f6        g3
4 f7        g2
5 f8=Q        g1=Q
6 Qc8+        Ka5
If 6 ... Kb6 7 Qb7+, Kc5 (on 7 ... Ka5 8 Qb5#.) 8 Qa7+ and White wins the Queen.
7 Qc3+        Kb6
If 7 ... Ka6 8 Qc4+, Kb6 9 a5+, Kxa5 10 Qb4+, Ka6 11 Qa4+, Kb6 12 Qa7+, and White wins the Queen.
8 a5+        Kb5
9 Qb4+        Ka6
On 9 ... Kc6 10 Qb7+, Kd6 11 Qb6+ is a brutal but convincing exchange of a Queen.
10 Qb7+        Kxa5
11 Qb4+        Ka6
12 Qa4+        Kb6
13 Qa7+
White wins the Queen and the game.
###79
1  f3
FORCING an exchange will leave White with a passed Pawn.
    1 ...         gxf3+
If Black avoids the exchange by 1 ... g3, then 2 f4 followed by 3 Kf3 wins the luckless Pawn.
2 Kxf3        Kf5
3 g4+        Kg5
4 b4
In order to block the position on the Queen side. Black's Pawns will be easier to pick off if they are immovable.
4 ...        b6
5 a4        Kg5
6 Kf4        Kf6
7 g5+        Kg6
8 b5        a5
9 Kg4        Kg7
10 Kf5        Kf7
11 Ke5
Poses a difficult problem. How does Black protect his Pawns while contending with the dangerous passed Pawn?
11 ...        Ke7
12 g6        Ke8
13 Kd6        Kf8
14 Kc6
White wins
###80
UNASSISTED by the King, White's 
Pawns break through by sheer force.
1  c5                   Kg4
Any Pawn move instead loses immediately: 1 ... dxc5 2 bxc5 followed by 3 d6 wins, or if 1... b6 2 axb6 gives White a passed Pawn.
2 b5!
Threatens to continue with 3 c6, bxc6 4 bxa6 and White wins.
2  ...    dxc5
If 2 ... axb5 3 c6 is decisive.
3 b6        cxb6
4 d6
White wins, his Queen Pawn having a clear road to the last square.
###81
1 c5!
THE right moment for the break-through. Black must not be given time for 1 ... Kc8, consolidating his position.
1 ...    bxc5
If 1 ... dxc5 2 d6, cxd6 3 Kxd6 b5 followed by 4 Kc6, and all Black's Pawns will fall.
2 Kb5    Kd7
3 a4
Not at once 3 Kxa5, as after 3 ... c6 Black has counter-play.
3 ...    Kc8
Whereas if now 3 ... c6+ (instead of 3 ... Kc8) 4 dxc6, Kc7 5 b3, and Black is helpless.
4 Kxa5        Kb7
5 Kb5        Ka7
6 Kc6        Kb8
7 a5        Kc8
8 a6        Kb8
9 a7+        Kxa7
The rest is no strain on White.
10 Kxc7           Ka6
11 Kxd6           Kb5
12 b3                Kb4
13 Kc6

White wins
###82
A FLOCK of Pawns can make life miserable for a Queen, however nimble she may be. 

1 a6      f1=Q
2 a7
Threatens to advance, become a Queen and checkmate.
2 ...           Qa1
If the King tries to flee by 2 ...   Kg8, then 3 a8=Q+, Kf7 4 Qb7+, Kxf6 (on 4 ...Kf8 5 Qf7+, Kg8 6 Qe8#) 5 Qg7+, Kf5 6 Qf7+, and Black's Queen falls.
3 f7        Qa3
Prevents the Pawns on the seventh taking another step.
4 d6
Poses a problem, since the Queen Pawn (which may not be taken) screens the Bishop Pawn, which threatens to Queen.
4 ...        Qf3
Solves the problem—temporarily.
5 d5!
How now?
5 ...        Qxf7
6 a8=Q+    Qg8
7 Qa1+        Qg7+
8 Qxg7#
White wins
###83
BLACK is a Pawn ahead, and has two connected passed Pawns. His advantage in Force, Space and Time is enough, according to Znosko-Borovsky, to guarantee a win. Despite his superiority in these elements Black is lost! White wins this by an intangible element—Strength of Position!
The fact that victory is brought about by the doubled Pawns, generally a weakness, is an amusing feature.
1 f5
Threatens to continue with 2 e6, breaking through for a Queen.
1 ...         e6
If 1... gxf5 instead, 2 e6, fxe6 3 g6 does the trick.
2 fxe6        fxe6
3 f4        Kb8
4 f5!        exf5
Or 4 ... gxf5 5 g6, and White wins.
5 e6        Kc8
6 e7
White wins
###84
WHITE breaks through what seems an impregnable barrier.
1 g3+    fxg3+
Against a refusal to capture, White proceeds as follows: 1 ... Kg5 2 g4, Kh4 3 Kg2, Kg5 4 Kh3, Kh6 5 Kh4, Kg6 6 g5, Kg7 7 Kh5, Kh7 8 g6+, Kg7 9 Kg5, Kg8 10 Kf6, Kf8 11 Ke6, Kg7 12 Kd7, and White wins.
2 Kg2        Kh5
3 Kxg3        Kg5
4 f4+        exf4+
5 Kf3        Kg6
6 Kxf4        Kf6
7 e5+        dxe5+
8 Ke4        Kf7
9 Kxe5        Ke7
10 d6+        cxd6+
11 Kd5        Ke8
12 Kxd6        Kd8
13 c7+         Kc8
14 Ke6
Unforgivable would be the precipitate 14 Kc6, allowing a draw by stalemate.
14 ...        Kxc7
15 Ke7        Kb8
16 Kd6        Kb7
17 Kd7        Kb8
18 Kc6        Ka7
19 Kc7        Ka8
20 Kxb6        Kb8
21 Ka6        Ka8
On 21 ... Kc7 22 Ka7 escorts the Pawn to the eighth square.

22 b6    Kb8
23 b7    Kc7
24 Ka7
White wins
###85
IT IS clear that White's hopes of winning depend on Queening the Rook Pawn, but how does he start? If 1 a4 (the natural move) Kg3!  2 a5, h5 3 a6, h4 4 a7, h3 5 a8=Q, h2#! Or if 1 Kxg2, Kg5 2 a4, bxa3e.p.  3 bxa3, Kf6 4 a4, Ke7 5 a5, Kd8 6 a6, Kc8 7 a7, Kb7, and Black wins the Pawn and the ending.

1 f6!
Throws a Pawn in the path, so that Black's King cannot reach d8 directly 
1 ...         gxf6
2 Kxg2     Kg5
3 a4         bxa3e.p.
4 bxa3         Kf5
5 a4         Ke5

Now if 6 a5, simply 6 ... Kxd5, while 6 c6 is met by 6 ... d6 followed by 7 ... Kxc6, wining for Black. Therefore:
6 d6!        cxd6
Clearly, if 6 ... c6 7 a5, Kd5 8 a6 wins.
7 c6!
White strews the Pawns about, and Black stops to pick them up, as Atalanta did the golden apples that Milanion threw in her path in their legendary race.
7 ...        dxc6
8 a5
The Pawn, safely out of reach, is sure to reach the Queening square.
White wins
###86
1 h6        Nd6
2 h7        Nf7+
3 Ke7        Nh8
IF 3 ... Ne5 instead (playing for 4 h8=Q, Ng6+ winning the Queen) 4 Kf6, Nd7+ 5 Kg7 keeps the Knight at a distance, and wins.
If White is careless now in chasing the Knight, this may happen: 4 Kf8, Ke5 5 Kg7, Ke6 6 Kxh8, Kf7, and White has been stalemated.
4 Kf6!
Maintains the opposition. White will continue with 5 Kg7, capture the Knight and win.
###87
1 e6!        Ne2+
WHITE must move out of check, and has choice of seven squares. One, and only one, is the right square that assures White of a win. Moving to any of the others leads to a draw, as follows:
2 Kf3, Nd4+, winning the Pawn
2 Kh3, Nf4+, winning the Pawn
C] 2 Kg4, Nc3 3 e7, Nd5 4 e8=Q,
Nf6+, winning the Queen
2 Kh4, Nf4     3 e7, Ng6+, winning the Pawn
2 Kf2, Nc3    3 e7, Ne4+,   4 Ke3, Nd6, stopping the Pawn (and capturing it later with the King).
2 Kg2, Nf4+, winning the Pawn
How that Knight hops around!
2 Kh2!!
This is the right move, to which there is no reply.
White wins
###88
1 Kd5
IF AT once 1 b7, Ne5+ 2 Kd5, Nd7 3 Kd6, Nb8 4 Kc7, Na6+ 5 Kb6, Nb8 6 Ka7, Nd7, and Black draws.
1 ...    Ne5
The best chance. There is no hope in 1... Ne3+ 2 Kc5, nor in 1... Nf6+ 2 Kc6. Now comes an amusing continuation:
2 g3+    Kf5
3 g4+    Kf6
4 g5+    Kf5
5 g6    Kf6
The King is torn between obligation to his Knight, and the necessity (let alone desire) to capture the impudent Pawn.
6 g7
White wins, having set Black a task too much even for a King to cope with.
###89
CLEVER defense disposes of one White Pawn, but the one that remains renders Black's Knight hors de combat.
1 g6        Nd6
2 g7        Ne8
3 g8=Q        Nf6+
4 Kg5!        Nxg8
5 d6!
Prevents the Knight from emerging, and threatens to move on and Queen with check.
5 ...        g3
All that is left, neither King nor Knight being able to head off the Queen Pawn.
6 d7    g2
7 d8=Q+
White wins
###90
1 Kg6        Ne5+
2 Kf6        Ng4+
3 Ke6        Nxh6
4 b6            Nf7
BLACK is willing to sacrifice one Knight so that the other can overtake the dangerous advanced Pawn.
5 Kxf7        Ne45
Attacks both Pawns.
6 b7        Nd6+
Did White overlook this?
7 Ke7        Nxb7
8 b4
Paralyzes the Knight.
8 ...        Kg5
9 Kd7        Kf6
10 Kc7
White captures the Knight and wins
###91
1 a5        Bf8
IN ORDER to stop the Pawn, Black tries to get his Bishop to the diagonal running from his a7 to g1.
2 Kd5
White of course plays to keep the Bishop off the line.
2 ...        Bh6
Aiming at the square e3.
3 g5+!        Bxg5
If 3 ... Kxg5 4 a6 followed by 5 a7 and the coronation at a8.
4 Ke4!        Bh4
5 Kf3!
White wins. His Pawn will march on, fearing neither the King who is too far off, nor the Bishop who can no longer interfere with its progress.
###92
1 Kd5!
WHITE does not play 1 b7 immediately, as after 1 ... Be5 in reply followed by 2 ... Bb8, the position is a draw.
1 ...        Be5
2 g3+        Kf5
3 g4+        Kf6
The alternative is 3... Kf4 leading to 4 g5, Kf5 5 g6, Kf6 6 g7, and White wins.
4 g5+        Kf5
5 g6        Kf6
If 5 ... Bc3 instead, then 6 b7, Be5 7 b8=Q, Bxb8 8 g7 wins.
6 g7 
White wins
###93
1 Kf5
BEFORE advancing the passed Pawn, White fights to keep the Bishop from occupying the long diagonal.
1 ...        Bb6
2 Ke4
But not 2 Ke5, as after 2 ... Ke2 3 h7, Kd3, and the Bishop reaches d4.
2 ...        Bd8
3 Ke5!        Bg5
If Black tries 3 ... Be7, then
4 h7, Bxb4 5 Kd4, Ba3 6 Kc3, and control of the long diagonal wins for White.
4 h7        Bc1
5 Kd5        Bb2
6 Kc5        Kf2
7 Kxb5        Kf3
Since the Bishop will not be able to cope with both passed Pawns, Black rushes his King over to dispose of the Rook Pawn.
8 Kc6        Kf4
9 b5        Kf5
10 b6        Kg6
Is Black in time?
11 b7        Be5
12 b8=Q        Bxb8
13 h8=Q
White wins
###94
1 Kg7        Bb3
2 h5        Kd7
APPARENTLY Black will attend to the Rook Pawn with his Bishop, while his King blockades the King Pawn—but White has a few surprises!
3 h6        Bc2
4 Kf7!
Threatens to win by 5 e6+ followed by 6 e7. The natural move 4 Kf6 (instead of 4 Kf7) allows Black to reply 4... Ke8 and draw the position.
4 ...        Bb3+
5 e6+!
Astonishing! Not only does White give away a Pawn, but he lets it be captured with check!-a beauty!
5 ...        Bxe6+
6 Kf6        Bg8
7 Kg7
White wins
###95
PASSED Pawns on both sides of the board can make life difficult for a lone Bishop.
1 c5        Kg5
2 c6        Bd8
3 Ke5        Kh6
To meet 4 Ke6 with 4 ... Kg7  5 f6+, Kf8, and Black has a draw.
4 f6!        Kh7
5 Ke6        Kg8
6 Kd7        Ba5
7 Ke8
White wins
###96
1 a5!
THE plausible 1 b5 would be refuted by 1 ... Kd8 2 Kb7, Kd7 3 b6, Ba5! 4 Ka7, Bxb6+ 5 Kxb6, Kc8 6 Ka7 (or 6 a5, Kb8) Kc7 7 a5, Kc8 8 a6, Kc7, and Black draws.
1 ...        Kd8
If 1 ... Bxb4 2 a6, and the Pawn cannot be stopped.

2 a6        Bf2
3 Kb7    Kd7
4 b5        Kd8
5 a7        Bxa7
6 Kxa7
White wins
###97
1 c7        Bc6+
2 b7!
UNEXPECTED, but the move that wins. If instead 2 Kb8, Bb7 3 b5+, Kxb5, 4 c8=Q, Bxc8 5 Kxc8, Kxb6, and Black draws. Or if 2 Kb8, Bb7 3 b8=Q, Bxc8 4 Kc7, Kb5 (but not 4 ... Bb7 5 b5+, Kxb5 6 Kxb7, and Black has been swindled) 5 Kxc8, Kxb6 and the position is a draw.

2 ...        Bxb7+
3 Kb8        Kb6
4 b5!
White wins, as Black must either abandon his Bishop, or move it and allow the Pawn to Queen.
###98
EVEN without the help of the King, a couple of passed Pawns can be re-markably effective against a Bishop.

1 f5        Kg3
2 g5!        hxg5
If 2 ... Bxh5 3 gxh6 wins at once, or if 2 ... Kg4 3 g6, Bd5 4 f6, Kxh5 5 f7, and White wins.
3 h6        Bg8
4 f6        Kf4
5 h7        Bxh7
6 f7
White wins
###99
1 Kg5

THE winning idea is to bring the
King around to d8.
1 ...         Kf7
2 Kf5        Bd7+
3 Ke5         Ba4
4 Kd6        Bb5
5 Kc7        Ba4
6 Kd8         Bb5
7 g8=Q+          Kxg8
8 e8=Q+         Bxe8Q
9 Kxe8
White wins
###100
1 c6
THE impulsive 1 a7 is met by 1 ... Bg2, wrecking White's chances of a win.
1 ...        dxc6
2 Kc5
Here too if 2 a7, c5+ 3 Kxc5, Bg2, and Black draws.
2 ...        Bc8
3 a7        Bb7
4 Kd6
Threatens to win by 5 Ke7 followed by 6 f7+.
4 ...        Kf7
5 Kc7        Ba8
6 Kb8
White wins. He captures the imprisoned Bishop and then Queens his Rook Pawn.
###101
1 c7        Rd6+
2 Kb5
EVERY move of White's must be timed right. For instance, if 2 Kb7, Rd7 followed by capturing the pinned Pawn draws, or if 2 Kc5, Rd1 followed by 3 ... Rc1+ with a Pawn skewer capture, does likewise.
2 ...        Rd5+
3 Kb4        Rd4+
4 Kb3        Rd3+
5 Kc2        Rd4
With this idea: 6 c8=Q, Rc4+! 7 Qxc4, and Black draws by stalemate.
6 c8=R!
Under-promotes to a Rook, and foils the stalemate try. Pieces are now even, but White threatens to mate by 7 Ra8+.
6 ...        Ra4
Kb3
Attacks the Rook and threatens 8 Rc1 mate. White wins a Rook and the game.
###102
1 Kf7        Rf2+
OBVIOUSLY if 1 ... Ra8 2 g8=Q, Rxg8 3 Kxg8, and the remaining Pawn becomes a Queen.
2 Ke6        Re2+
The idea is to annoy the King enough to make him come down the board, whereupon the Rook can get to the first rank and hold back the Pawns.
3 Kf5     Rf2+
If at once 3 ... Re8, then 4 Kf6, Kc7 5 Kf7, Kd7 6 g8=Q wins for White.
4 Ke4        Re2+
5 Kf4!
The natural 5 Kf3 fails after 5 ... Re8 6 Kf4, Rg8 7 Kf5, Rxg7, and the Rook giving up his life for the second Pawn, draws for Black.
5 ...    Re8
6 Kg5    Kc7
7 Kh6    Kd7
8 Kh7    Ke7
9 g8=Q
White wins
###103
A LONE Knight can mate! Without a Pawn on the board, Black would be quite safe. As it stands, the Pawn is compelled to move down and hem in its own King.
1 Kcl        a3
2 Nc2+        Ka2
3 Nd4        Ka1
4 Kc2        Ka2
Or 4 ... a2 5 Nb3#.

5 Ne2        Ka1
6 Nc1        a2
7 Nb3#
###104
1 Nc5        Kd2
2 Na4!
IF WHITE had played 2 Ne4+ instead, the reply 2 ... Ke3 would threaten to capture the Knight and then the Pawn. After White's actual move, Black must lose time chasing after the Knight.
2 ...         Kd3
3 Kg2         Kc4
4 Kf3         Kb3
5 Ke4         Kxa4
6 Kd5
White has gained time to bring his King to a dominating position. The rest is elementary.
6 ...         Kb5
7 c4+        Kb6
8 Kd6        Kb7
9 c5         Kc8
10 Kc6!
White wins
###105
1 Kd3
WHITE does not go after the Pawn, as 1 Kb4, Ka1 2 Kxa3 is a draw by stalemate.
1 ...         Ka1
2 Na4!
Instead, he gives up his own valuable Pawn!
2 ...        Kb1
3 Kd2        Ka1
4 Kc1
Forces Black to take it.
4 ...         Kxa2
5 Kc2         Ka1
6 Nc5         Ka2
7 Nd3         Ka1
8 Nc1         a2
9 Nb3#
###106
1 Na2!
IMMOBILIZES the Pawn, and avoids such dangers as this: 1 Kf5, Kg8 2 Kg6, Kh8 3 h7, a2 4 Nxa2, stalemate.
1 ...        Kf8
2 Kf6!        Kg8
3 Kg6        Kh8
4 Nb4        Kg8
5 h7+        Kh8
6 Nc6        a2
7 Nd8        a1=Q
8 Nf7#
###107
1 h3
ON 1 Ng5+, Kg4 draws easily.
1 ...        Kg3
2  Ng5        Kf4
3  Ne4        Kf3
4 Kd4        Kf4
There is no fight in 4 ... Kg2 5 Ng5, Kg3 6 Ke3, Kg2 7 Kf4, Kf2 8 Kg4, and White wins.
5 Kd5         Kf5
Still keeping White's King at bay.
6 Nc3!
Very subtle! The natural 6 Nf2 leads to 6 ... Kf4 7 Ke6, Kg3 8 Kf5, Kxf2 9 Kg4, Ke3 10 Kxh4, Kf4 11 Kh5, Kf5 12 Kh6, Kf6 13 Kh7, Kf7 14 h4, Kf8 15 h5, Kf7 16 h6, Kf8, and White must either play 17 Kg6 allowing Black to reach g8 and a draw, or imprison his own King in a humiliating stalemate position.
6 ...        Kf4
7 Ne2+        Kf3
8 Ng1+        Kg2
9 Ke4        Kxg1
10 Kf3        Kh2
11 Kg4        Kg2
12 Kxh4        Kf3
13 Kg5        Ke4
14 h4            Ke5
15 h5            Ke6
16 Kg6
White wins, as the Pawn is free and Black can never reach g8, the drawing square.
###108
IT WOULD appear that Black has at least a draw, since White cannot hold on to his solitary Pawn—but White has hidden resources!
1 Ng2!        hxg3+
The alternative is 1 ... Kh2 2 Nxh4, Kh3 3 Nf5, h5 4 Kgl, h4 (unfortunately forced) 5 gxh4, and the Rook Pawn marches up the board to Queen.
2 Kg1        h5
3 Kh1        h4
4 Nf4#
###109
1 Kd7!
BLOCKING the Pawn—a strange winning move! White does not advance the Pawn immediately, as after 1 d7, Bc6 followed by 2 ... Bxd7 draws the position.
1 ...         Kd5
2 Kc7
Now threatening 3 d7.
2 ...        Bc6
3 Ne4!
Black is in Zugzwang. He loses because he is compelled to move.  He has these choices:
A]    3 ... Kxe4 4 Kxc6, and the Pawn Queens.
B]    3 ... Be8 4 Nf6+ winning the Bishop.
C]     3 ... Bb5 4 Nc3+, Kc5 5 Nxb5, Kxb5 6 d7, winning.
D]     3 ... Ba4 4 Nc3+ and 5 Nxa4.
White wins
###110
1 Nf5        Ba8
OBVIOUSLY, if 1 ... Bc8 2 Ne7+ followed by 3 Nxc8 wins for White.
2 Nd4+        Kc5
3 Ne6+        Kc6
The only move, since the King must stay near the Pawn, and 3 ...  Kb5 loses by 4 Nc7+.
4 Nc7         Bb7
5 Nd5!
Leaves Black without resource. On   5 ... Kxd5 (or any other move by the King) the reply 6 Kxb7 wins for White. Should Black try 5 ... Bc8, then 6 Ne7+, Kd7 7 Nxc8, Kxc8 8 b7+ compels resignation and
White wins
###111
1 Nc6
SHUTS off the Bishop's action on the long diagonal so that the Pawn may march through. The Pawn now threatens to reach b7 safely.
1 ...        Bf1
2 b6        Ba6
3 Kd6        Bb7
If the King tries to help by 3 ... Ke3, then 4 Kc7, Ke4 5 Nb4 evicts the Bishop, enabling the Pawn to advance next move.
4 Kc7        Ba8
5 Nd8        Ke3
On 5 ... Bf3 6 Nc6 shuts the Bishop out permanently, assuring the Pawn of Queening in two more moves.
6 Nb7
Brutal but necessary. The Bishop is imprisoned.
6 ...         Kd4
7 Kb8        Kd5
Unfortunately for Black, his King may not move to c5 attacking the Pawn.

8 Kxa8         Kc6
9 Ka7
White wins
###112
WHITE plans to maneuver his Knight to b7, shutting off the Bishop, and then promote his Pawn. White's object is clear, but he must time his moves right to accomplish it.

1 Na5
The plausible try 1 Kb8, Kd8 2 Na5 (threatening 3 Nb7+ followed by advancing the Pawn) fails after 2 ... Ba8! 3 Kxa8, Kc7 4 Nc4, Kc8 5 Nd6+, Kc7 6 Nb5+, Kc8 due to the fact that the Knight cannot gain a move.
1 ...        Ba8

White was threatening to shut off the Bishop by 2 Nb7 or 2 Nb6+.

2 Kc8!
Gaining a move for the final maneuver—something the Knight cannot do. If at once 2 Kb8, Kd8 3 Nb7+, Kd7 4 Kxa8, Kc8 5 Nd6+, Kc7 6 Ne8+, Kc8, and the stubborn King cannot be ousted.
2 ...        Kd6
3 Kb8        Kd7
4 Nb7
Black's King needs a waiting move (so that 5 Kxa8 can be met with
5 ... Kc8) but he has none!
4 ...        Kc6
5 Kxa8        Kc7
6 Nd6!
Prevents Black from moving to c8, and enables his own King to emerge.
6 ...        Kxd6
7 Kb7
White wins
###113
AN INTERESTING duel between Knight and Bishop.
1 Nb7+        Kb4
The other defenses lead to the following play:
A] 1 ... Kc6 2 Nd8+, Kb5 (on 2 ... Kb6, or Kd5 3 Ne6 followed by Queening with check wins, while 2 ... Kc5 3 Ne6+ costs the Bishop) 3 Ne6, Bg3 4 Kf6, Bh4+ 5 Ng5 wins.
B] 1 ... Kc4 2 Nd6+, Kd5 3 Kf6! Bg3 4 Kg5 wins. A beautiful variation.
C] 1 ... Kb6, d5 or d4 2 d8=Q+ wins.
2 Nd6        Be3
On 2 ... Bg5+ 3 Ke8, and White will then play 4 Nf5 followed by 5 Ne7, blocking the Bishop's action on the diagonal.
Black's actual move leaves the possibility of checking at g5 open, as well as the opportunity of shifting to another diagonal via b6.
3 Nc8        Bf4
Kept out of b6, the Bishop heads for c7.
4 Nb6!
Refutes 4 ... Bc7 by 5 Nd5+ winning the Bishop, and 4 ...  Bg5+ by 5 Ke8, Kb5 6 Nc8 followed by 7 Ne7 shutting out the Bishop.
4 ...         Kc5
After all the fencing, the position is almost identical with the opening scene.
5 Nd5!
White wins. If 5 ... Kxd5 6 d8=Q+, (emphasize Pawn Queens with check), while on 5 ... Bg5+ 6 Nf6 is decisive.
###114
A PRETTY tactical device in Pawn endings is the forcible displacement of a piece that prevents your Pawn from moving up the board.
A simple and pleasing example:
1 Ng7+!    Nxg7
On 1 ... Kf7 instead, 2 Nxf5, Kf6 3 h6 wins for White. The Pawn and Knight stay where they are until the King comes up and helps.
2 h6                 Kf7
3 h7
Certainly not the greedy 3 hxg7, when 3 ... Kxg7 drawing is fit retribution. 
White wins
###115
BEFORE the Pawn can advance, Black's Knight must be driven off and kept out of active play.
1 Nd2    Kg7
2 Nc4
Forces Black's Knight away from the center of action.
2...                    Nb1
On 2 ... Nc2 3 b5, Ne1 4 b6, Nd3+ 5 Kb5, Nf4 6 b7 leaves Black helpless.
3 Kd4!
Not at once 3 b5 as 3 ... Nc3 attacks the Pawn, and after 4 b6, Na4+ 5 Kc6, Nxb6 captures it and draws.
3 ...    Kf7
For the moment Black's Knight is out of business.
4 b5    Ke7
5 b6    Kd7
6 Kc5    Nc3
7 Ne5+        Kc8
Black has no time to attack the Knight as after 7 ... Ke6 8 b7, Na4+ 9 Kb4 wins.
8 Kc6        Ne2
9 b7+        Kb8
10 Kb6
White mates on next move with Nd7.
###116
1 Ka7!
A SUBTLE winning move. The natural 1 Kb8 leads to 1 ... Kb5 2 Nb4 (2 a7, Nc6+ wins the Pawn) Nc6+ 3 Kb7 (3 Kc7, Nxb4 4 a7, Nd5+ followed by 5 ...  Nb6 draws) Na5+ 4 Ka7, Kxb4 5Kb6, Nc4+ 6 Kc7, Ka5 7 a7, Nb6 and Black draws.
1 ...    Kb5
If 1 ... Nc6+ 2 Kb6, Kd5   (or 2 ... Ne7 3 Kb7) 3 Nb4+, Nxb4 4 a7 and wins, or if 1 ... Kb5 2 Nb4! (a pretty sacrifice to prevent the King coming closer) Kxb4 3 Kb6, and White wins.
2 Nb4!    Ka5
If 2 ... Kxb4 3 Kb4 wins, or if 2 ... Kc5 3 Kb8, Kxb4 4 Kc7 (certainly not 4 a7, Nc6+ followed by 5... Nxa7) Ne6+ 5 Kb6, and the Pawn cannot be stopped.
3 Kb8        Nc6+
4 Kb7        Nd8+
5 Kc7        Ne6+
6 Kb8        Nc5
On 6 ... Kb6 7 a7, Nc7 8 Nd5+ wins.
7 a7        Nd7+
8 Kc7        Nb6
9 Kb7        Kb5
10 Nd5
White wins. Black's Knight must not move and dares not stay. A full-bodied instructive Réti composition.
###117
WHITE must draw Black's Knight away from its post, so that his Pawn may advance and Queen.
1 Ne6
Intending to attack the Knight by 2 Nf8.
1 ...    Kd5
Vacates c6 for later occupation by his Knight.
2 Nf8        Ne5
3 Ka8
Clearly not 3 b8=Q, Nc6+, and the newly-crowned Queen comes off the board.
3 ...         Nc6
4 Nd7        Ke6
If 4... Kd6 5 Nb6, Kc7 (on 5... Kc5 6 Nc8 followed by 7 Na7 wins) 6 Nd5+, Kd6 7 Nb4 wins for White.
5 Nb6        Kd6
6 Nc8+        Kc7
7 Na7        Nb8
8 Nb5+        Kb6
9 Kxb8
White wins
###118
1 Nf6
TO PREVENT a perpetual check draw by 1 ... Nd7+ 2 Ka8, Nb6+ 3 Kb8, Nd7+ 4 Kc8, Nb6+ etc.
1 ...                    Na8!
A clever defense. If Black had played 1 ... Kc5, then 2 Kb7, Kb5 3 Nd5 evicts the Knight (from holding back the Pawn) and wins.
2 Nd5!
A fine reply. White avoids 2 Kxa8, Kc7 3 Nd5+, Kc8 4 Ne7+, Kc7, and Black draws, the Knight being unable to gain a move to dislodge Black's King.

2 ...        Kd7
Ready to meet 3 Kxa8 with 3 ... Kc8, and Black draws.

3 Kb7
Restricts Black to one move by his King, since his Knight may not come out.
3 ...        Kd8
The only move to keep in touch with the square c8. On 3 ... Kd6 instead, White replies 4 Kxa8 fol-lowed by 5 Kb8, winning.

4 Nb6        Nc7
5 Kc6

White wins
Black must either move his King, losing the Knight, or move the Knight and allow White's Pawn to Queen.
###119
A REWARDING ending where the variations are as brilliant as the main play.
1 Kg5!
Moving in on the Knight. Straightforward attack by 1 Kf5 leads to this: 1 ... Ka7! 2 Kf6 (or 2 Nd7, Nxd7 3 e6, Nb6 4 e7, Nc8 5 e8=Q, Nd6+ and Black wins the Queen) Kb6 3 Kg7, Kxc5 4 Kxf8, Kd5, and Black captures the Pawn and draws. Or   1 Kf5, Ka7 2 Kg5, Kb6 3 Nd7+, Nxd7 4 e6, Nc5 5 e7, Ne6+ 6 Kf6, Nc7, and the Pawn can never move on to Queen.
1 ...               Ka7
2 Kf5!
The natural attack 2 Kh6 is met by 2 ... Kb6 3 Nd7+ (if 3 Kg7, Kxc5 draws) Nxd7 4 e6, Nf6 5 Kg6 (5 e7, Ng8+) Nd5 and Black draws.

2 ...        Kb6
If Black returns instead by 2 ... Kb8, then 3 Kf6, Kc8 (or 3 ... Nh7+ 4 Kg6, Nf8+ 5 Kg7, and the Knight is magnificently trapped) 4 Kg7, Nd7 5 Nxd7, Kxd7 6 Kf7, and the Pawn will Queen.
3 Nd7+!    Nxd7
4 e6        Nc5
If 4 ... Kb7 5 e7 followed by 6 e8=Q wins.

5 e7        Nb7
A last flicker of hope. If White precipitately Queens his Pawn then
6 ... Nd6+ removes the Queen.
6 Ke5
White wins
###120
THERE is classic economy in this pleasing example of persuasion exerted on Black's Knight to abandon his important post.
1 f6        Nf2+
2 Kg2        Nd3
3 f7        Nf4+
4 Kh2
In order later to restrict the moves of Black's King.
4...        Ng6
The Pawn is stopped, but White has a counter-combination.
5 Nf3+        Kg4
6 Ne5+!        Nxe5
7 f8=Q
White wins
###121
GALLOPING furiously after the passed Pawn, Black's Knight heads it off just in time—or does he?
1 f5        Ne3
2 f6        Nf5
3 f7        Ne7+
4 Kb7        Ng6
The Pawn is stopped, but now comes... .
5 Nxe5
Ready to reply to 5... Nxe5 with 6 f8=Q and White wins, or if Black plays 5 ... a3, then the continuation is 6 Nxg6, a7 7 f8=Q, a1=Q, 8 Qa8+ will remove Black's Queen.
5 ...        Nf8
6 Nc6#!
###122
WHITE'S first move is obvious, but how will his King hide from all the checks?

1 bxa7        Rb3+
2 Kc7        Rc3+
3 Kd7
The King must not leave the sev-enth rank. If for example 3 Kd6, Rc8 solves Black's problems, or if 3 Kd8, Rxh3, and White does not dare Queen his Pawn.
3 ...        Rd3+
4 Ke7        Re3+
5 Kf7        Rf3+
6 Kg7        Rg3+
7 Ng5!
Chess players are generous!
7 ...     .   Rxg5+
Now the Rook is closer, which is the purpose of White's sacrifice.

8 Kf7        Rf5+
9 Ke7        Re5+
10 Kd7        Rd5+
11 Kc7        Rc5+
12 Kb7        Rb5+
13 Kc6!
White wins. The Rook has no more checks, and cannot return to the first rank.
###123
WHITE is a whole piece ahead, but how does he win? Moving closer with his King stalemates Black, and meaningless Knight moves do not force the win.
Clearly, giving up the Knight will break through, but the sacrifice must be timed right. If for example 1 Nxh6, g7xN 2 Kf6, Kg8 3 g7, Kh7 4 g8=Q+ (4 Kf7 stalemates Black) Kxg8 5 Kg6, and White can win the remaining Pawn but not the ending.
1 Nd6    Kg8
2 Ne8    Kh8
3 Nf6!
Leaves no choice.
3 ...            gxf6
4 Kf7            f5
5 g7+            Kh7
6 g8=Q#
###124
HOW does White win? He can Queen his advanced Pawn, but so can Black.
1 Nb4+!        axb4
2 e7        d2
3 e8=Q        d1=Q
4 Qc6+        Ka5
5 Kb7!
Threatens mate on the move. 
5 ...        Qd3
The alternative 5... b3 is answered by 6 a3, and Black is faced with an additional threat of mate in two by 7 Qc5+, Ka4 8 Qb4#.
6 Qb6+         Ka4
7 Qa7+        Kb5
8 Qa6+
White wins the Queen and the game.
###125
1 b7             Nd6+
2 Kd4!
NOT at once 2 Kd5, as after 2 ... Nxb7 3 e7, Kf7, Black has a draw.
2 ...        Nxb7
3 Kd5
The difference is that Black is now confined to two plausible defenses.
3 ...        Kg7
Moving away from the Pawn, but if he makes the only move open to his Knight then this happens: 3 ... Nc5 4 e7, Na6 5 Kd6 (of course not 5 e8=Q, Nc7+ winning the Queen) Kf7 6 Nd8+, Ke8 7 Ne6, Kf7 8 Ng7, Nc7 9 Kd7, Kf6 10 Ne8+, Nxe8 11 Kxc7 and White wins.
4 Nd8!        Nxd8
5 e7
White wins
###126
BLACK is so busy warding off at-tacks on his Knight, and coping with the passed Pawns, that he overlooks a greater danger—that of being checkmated!
      1 Ke7        Nh7
Black can try to capture the Pawns, in which case the play would run:  1 ... Kg5 2 Kxf8, Kxh5 3 Kg7, Kg4 4 Ne4, Kf3 5 Kg6, Kg4 6 Kf6, Kf3 7 Kf5 and White wins.
    2 Kf7        Kg5
The alternative is 2 ... Nf6 when 3 h6, Kg5 (on 3 ... Ke5 4 Nd7+ wins) 4 Ne4+, Nxe4 5 h7 wins for White.
    3 Ne4+        Kxh5
    4 Kg7        Ng5
    5 Nf6#
###127
1 g6        Kf6
2 g7        Kf7
3 Nd5
THREATENS to win by 4 Nf6 followed by 5 g8=Q+.
3 ...        Ne6
Other defenses are:
A] 3 ... Nd7 (or h7) 4 Ne7, Nf6 5 g8=Q+, Nxg8 6 h7 and wins.
B] 3 ... Ng6 4 Nf6, Ne7 5 g8=Q+, Nxg8 6 h7 and wins.
4 Ne7        Nxg7
5 h7!
White wins
###128
ONLY two moves are needed to win, but they are enough to force Black into zugzwang. Zugzwang is the compulsion to move—and lose.
1 Nd4+     Kc5
On 1 ... Kb7 instead, the continuation is 2 Kxh2, Ka6 3 Nb3, Bf4+ 4 Kh3, Kb5 5 Kg4, Bb8 6 f4, Kb4 7 f5, Kxb3 8 f6, Kb4 9 f7, Bd6 10 a6, and White wins, the Bishop being outdistanced by the Pawns.
2 Kh1!
White wins! There are no threats, and if Black could pass nothing could happen to him. But zugzwang—the compulsion to move when it's one's turn to move—has him in its grasp, and he is lost. The proof: If
A] 2 ... Bf8 3 Ne6+ wins the Bishop.
B] 2 ... Bg7 3 Ne6+ wins the Bishop.
C] 2 . . Bg5 3 Ne6+ wins the Bishop.
D] 2 ... Bf4 3 Ne6+ wins the Bishop.
E] 2 ... Bd2 3 Nb3+ wins the Bishop.
F] 2 ... Bc1 3 Nb3+ wins the Bishop.
G] 2 ... Kd6 3 Nf5+ wins the Bishop.
H] 2 ... Kd5 3 a6, and the passed Pawn becomes a Queen.
###129
1 c6
BEFORE starting his Rook Pawn, White blocks the long diagonal.
1 ...             dxc6
2 a6          Bf3
If 2... c5 3 Ne5 prevents the Bishop from getting on the        diagonal.
3 Ng5          Bd5
4 Ne6          c5
Black may not bring his King closer, as after 4 ... Kd7 5 Nc5+, Kc7 6 a7 leaves him helpless.
5 Nc7+    Kd7
6 Nxd5    Kc8
On 6... Kc6 7 Kg2, c4 8 Kf2, c3 9 Ke3, and White overtakes the Pawn.7 Nb6+    Kb8
Forced, as 7... Kc7 allows 8 a7 followed by Queening.
8 Nd7+    Ka7
9 Nxc5
White wins. His King simply comes up the board and escorts the Pawn to the Queening square.
###130
1 c5
THREATENING to win by 2 c6, bxc6 3 b7.
1 ...         Bb1
2 Ne6!
A beautiful sacrifice as we will soon see.
Ready to circumvent 2 ... Be4 by 3 Ng5+, Kf4 4 Nxe4, Kxe4
5 c6, etc.
2 ...        fxe6
3 c6        Be4
4 c7
White wins
###131
1 f7        Rxa6+
2 Nf6        Ra8
3 Ne8        Ra6+
4 Kg5
THE King must find a way to evade the Rook's checks. He may not move to the seventh rank, as after 4 Kg7 for example, 4 ... Ra7 followed by sacrificing the Rook for the Pawn, draws. Nor may he move to the Bishop file, as after 4 Kf5, Ra1 followed by 5 ... Rf1 (with or without check) draws for Black.
4 ...         Ra5+
5 Kg4         Ra4+
6 Kg3         Ra3+
7 Kf2         Ra2+
8 Ke3         Ra3+
Now we go up the board.
9 Ke4        Ra4+
10 Ke5        Ra5+
11 Ke6        Ra6+
12 Kd7        Ra7+
13 Nc7
White wins, as the Rook has no more checks and cannot stop the Pawn.
###132
1 b7        Rf8
2 Nb4+        Ke4
BLACK plays to capture or render impotent the King Knight Pawn. Moving to the Queen side instead would lead to this: 2 ... Kc4 3 Nc6, Kc5 4 b8=Q, Rxb8 5 Nxb8, Kd5 6 g4, Ke5 7 g5, Kf5 8 Nd7, Kf4 9 Nf6 and White wins.
3 Nc6        Kf4
Hoping for this: 4 b8=Q, Rxb8 5 Nxb8, Kg3, and Black captures the Pawn next move and draws.
4 g4!        Kxg4
What else is there? Black has prevented 5 g5 or allowing 6 b8=Q+, and if he tries 4 ... Rf6+ (instead of 4 ... Kxg4) then 5 Kg7 (not 5 Kxh7 which is refuted by 5 ... Rf7+, and White loses the ambitious Queen Knight Pawn) Rg6+ 6 Kxh7 and White wins.
5 Kg7        Re8
The only square open along the rank.
6 Kf7        Rh8 
Once again the only square left.
7 Ke7!
Threatens 8 Nd8, cutting off the Rook, after which the Pawn could advance.
7 ...        Rg8
8 Nd8    Rg7+
If White should now get out of check by moving his King, the reply 9 ... Rxb7 would draw easily, the Knight alone being unable to force mate.
9 Nf7        Rg8
But now the Rook returns to the first rank. Does Black get a draw after all?
10 Nh6+
Certainly not! The Knight fork wins the Rook and the game.
###133
WHITE cannot save his stranded Knight, but he can close in and cap-ture the badly-placed Bishop in return. At the critical moment though, he must find the pretty surprise move that turns an apparent draw into a win.
1 Kg4        Kc8
2 Kh5        Kd8
3 Ng7!        Bxg7
4 h8=Q+!
Unexpected, and the only way to win. If instead 4 Kg6, Bh8 5 Kf7, Kd7 6 Kg8, Ke7 7 Kxh8, Kf7, and White is stalemated.
4 ...        Bxh8
5 Kg6        Ke7
6 Kh7        Kf7
7 Kxh8        Kf8
8 Kh7        Kf7
9 Kh6        Kf8
10 Kg6        Ke7
11 Kg7        Ke8
12 Kxf6        Kd8
13 Ke6        Ke8
14 f6        Kf8
15 f7        Kg7
16 Ke7 

White wins
###134
A LONE Bishop can mate—if the King is imprisoned by his own friends!
1 Kc7        a3
2 Ba4        a2
3 Kc6        a1=Q

The Pawn must push on and become a Queen (or some other piece) though the enemy is closing in on his King.
4 Bb5#
###135
BISHOP and Rook Pawn win easily against the lone King if the Bishop controls h8, the Queening square. If the Bishop is the wrong color, as in the diagrammed position, the win can only be attained if the adverse King is kept out of the corner square. Once he reaches that square, a win is impossible.
1 h6        Kf7
2 Bh7!
To prevent 2 ... Kg8 followed by 3 ... Kh8, which would assure Black of a draw.
2 ...        Kf6
3 Kf4        Kf7
4 Kf5        Kf8
5 Kf6        Ke8
6 Bf5        Kf8
7 h7        Ke8
8 h8=Q#
###136
AGAINST a Rook Pawn and a bad Bishop (one that does not control the Pawn's Queening square) Black can draw if his King can reach the corner square.
White's play to win is easily understood once we realize that his purpose is to keep Black's King from reaching g8 or g7, from whence he can move to h8 and the safety of a draw.
2 h6        Kf6
3 Bf5!
To prevent 3 ...  . Kg6 attacking the Pawn.
3 ...        Kf7
Now threatening to move to g8, and a sure draw.  Not so fast!
4 Bh7!        Kf6
5 Kf4        Kf7
6 Kf5        Kf8
7 Kf6        Ke8
8 Bf5        Kf8
White mates in two moves, e.g.
9 h7        Ke8
10 h8=Q#
###137
CAN White overtake the Rook Pawn?

1 Kf7        h5
2 Ke6
But not 2 Bf6 as 2 ... Kb3 3 Kg6, h4 4 Bxh4, Kxb2 and Black has a draw.
2 ...        h4
This time if Black had played 2 ... Kb3, 3 Kf5 would close in on the Pawn.
3 Kd5
Ready to meet 3 ... Kb3 with    4 Ke4 and the Rook Pawn is doomed.
3 ...        h3
But this certainly looks good!
4 Kc4!        h2
5 Bb4        h1=Q
6 b3#!
And not a moment too soon.
###138
1 a7
WHITE can try to prevent Black from Queening his Pawn, but after 1 Bc3, Kc8 2 Be5 (to stop 2 ...  Kb8 when Black has a certain draw) b2 3 Bxb2 Kb8, and White cannot possibly force a win.
1 ...        b2
2 a8=Q        b1=Q
With Black's King and Queen so far apart from each other, it seems incredible that the Queen will either be lost or the King driven into a cul-de-sac in a half-dozen moves.
3 Qb7+        Ke6
If 3 ... Ke8 4 Qe7#, or if 3 ... Kd8 4 Be7+, discovering an attack on the Queen.
4 Qe7+        Kd5
The alternative is 4 ... Kf5   5 Qh7+ winning the Queen.
5 Qd6+        Kc4
Here if 5 ... Ke4 6 Qg6+ wins the Queen.
6 Qc5+        Kb3
7 Qc3+        Ka2
8 Qa3#
###139
IT IS clear that Black can draw if his King reaches a8, since White has a Rook Pawn and his Bishop does not control the Pawn's Queening square. How does White keep the King out? If he plays the natural 1 Bc5, the continuation would be 1... Kf7 2 a4, Ke6 (but not 2... Ke8 3 a5, Kd8 4 Bd6, Kc8 5 a6, Kd8 6 a7 and White wins) 3 a5, Kd5 4 a6, Kc6 5 Kg2, Kc7 (aiming for the square b8 and a sure draw) 6 Ba7, Kc6 (threatens to win the Pawn) 7 Bb8, Kb6 8 a7, Kb7, and White must concede the draw.
1 Bb4!
A subtle move, whose purpose will shortly be manifest.
1 ...         Kf7
2 a4        Ke6
3 a5        Kd5
4 a6        Kc6
5 Ba5!
Prevents 5... Kc7, and closes all the roads leading to Black's a8.
5 ...         d5
6 Kg2        d4
7 Kf3        d3
8 Ke3        Kd7
The King has nothing left but retreat.
9 a7
White wins
###140
BEFORE Black can be subdued, his King must be prevented from reaching the corner square (a8), and his remaining Pawn captured.

1 Bc5        Ka5
2 Kb7        Kb5
3 Bb6!
Forces Black to move down the board.
3 ...         Kc4
4 Kc6        Kb3
Black can try a roundabout route to a8, with this result: 4 ... Kd3 5 Kb5, Ke4 6 Kxa4, Kd5 7 Kb5, Kd6 8 Ka6, Kc6 9 a4, Kd7 10 Kb7, and Black is kept out - with plenty of protection for a Pawn Queening.
5 Bc5        Kc4
6 Bd6        Kd4
7 Kb5        Kd5
8 Bh2        Ke6
9 Kxa4    Kd7
10 Kb5        Kc8
11 Kc6
White wins
###141
1 Bd4!
NOT at once 1 c8=Q, as Black replies by Queening with check.
1 ...     g1=Q
Black cannot play 1 ... Kxd4 as after 2 c8=Q, g1=Q 3 Qc5+ forfeits his newly-crowned Queen.
2 Bxg1    a2
Now this Pawn looks dangerous!
3 Bd4!
Obviously, if White had played 3 c8=Q instead, then 3 ... a1=Q would have drawn for Black.
3 ...         Kxd4
4 c8=Q        a1=Q
5 Qh8+
White wins the Queen and the game.
###142
WHITE must do something about the menacing Rook Pawn, but after the plausible 1 Bb1, h4 2 Bf5, d6 3 Bh3, Kd3 4 Bf1+, Kd4 5 Kb5, Kc3 6 Kc6, h3 7 Kd5 (or 7 Kxd6, h2 8 Bg2, Kxc4 and a draw) h2 8 Bg2, Kd3, he can capture the Queen Pawn only at the expense of his own Pawn. Therefore:
1 c5!        h4
2 Be6!
A bit of generosity saves miles of analysis.
2 ...        dxe6
3 c6        h3
4 c7        h2
5 c8=Q        h1=Q
6 Qc3+        Kd5
7 Qc5+        Ke4
8 Qc6+
White wins the Queen and the game.
###143
WHITE must retain his Pawn in order to win. He must prevent something like this happening: 1 ... c3 2 bxc3, Kc4, and his invaluable Pawn comes off the board, enabling Black to draw.
1 Bg1+        Kb4
2 Bd4        Kb3
3 Bc3        b4
4 Kd4!
But not 4 Bd4, c3, and Black forces the draw.
4 ...        b xc3
5 bxc3        Ka4
6 Kxc4
White wins
###144
ORDINARILY, Black could move his King to the corner and draw easily against a Rook Pawn and a Bishop of the wrong color. White wins the position though, by effecting a change of identity in his impotent Rook Pawn!
1 Kc6        Ka8
If Black plays 1 ... Ka6 instead, there comes 2 Be3, Ka5 3 Bc5, Ka6 4 Bb6, b4 5 axb4, a3 6 b5#.
2 Kb6!
White does not blunder into 2 Kxb5? after which there is no possible way to win.
2 ...        b4
3 axb4
But now the Rook Pawn has been transformed into a Knight Pawn!
3 ...        a3
4 b5        a2
5 Be5        a1=Q
6 Bxa1    Kb8
7 Be5+    Ka8
8 Kc7        Ka7
9 b6+
White wins
###145
BLACK has a wicked-looking Rook Pawn, but an alert sacrifice renders it harmless.
1 Bf3+        Kg1
2 Bh1!        Kxh1
3 Kf1        d5
4 exd5        e4
5 d6        e3
6 d7        e2+
7 Kxe2        Kg1
8 d8=Q        h1=Q
Black gets his Queen after all, but will never have a chance to move it.
9 Qd4+
Care must be taken in checking. Such a move as 9 Qg5+ for example allows Black to interpose his Queen with check, forcing a draw, i.e. 9 ...  Qg2+.
9 ...        Kh2
10 Qh4+        Kg2
On 10 ... Kg1 11 Qf2#.
11 Qg4+        Kh2
12 Kf2
White wins
###146
BLACK is a piece behind, but his passed Pawns on both sides of the board can be a worry to a Bishop.
1 Bb1!
Stops the dangerous-looking Rook Pawn from moving on.
1        f4
Or the following: 1 ... Kb3 2 Kc6, Kb2 3 d7, Kxb1 4 d8=Q, a2 5 Qb6+, Ka1 6 Qxa5, Kb2 7 Qb4+, Kc2 8 Qa3, Kb1 9 Qb3+, Ka1 10 Qc2, f4 11 Qc1#.

2 Kc 6        f3
3 Kc5!
Threatens mate on the move.
3 ...        Kb3
4 d7        f2
5 d8=Q        f1=Q
6 Qd5+        Kc3
Or 6... Kb2 7 Qa2+, Kc3 8 Qc2#.
7 Qd4+        Kb3
8 Qa4+!        Kxa4
If he refuses the Queen by 8 ... Kb2, then 9 Qc2+, Ka1 10 Qa2#.
9 Bc2#
###147
IN THIS battle of Bishop against Knight, White plays to force his opponent into zugzwang.

1 Bf7            Kb5
The King stays near the Pawn, preventing White from capturing the Knight.
2 Be8+           Ka5
3 Bd7
Not 3 Bh5 when Black draws by 3 ... Ne6+ 4 Kc6, Nd8+ 5 Kc5, Ne6+ 6 Kc6, Nd8+, and Black has a perpetual check- aka draw.
3 ...           Ka6
4 Bh3                Ka5
Black may not move the Knight as after 4 ... Nb7, White can win the beast either by 5 Bg2 or 5 Bf1+.
5 Bg4
Still keeping an eye out for the Knight check.
5 ...           Kb5
6 Be2+       Ka5
If 6 ... Kc5 instead, the reply 7 Bc4 leaves Black helpless.
7 Bc4
Now we have the diagrammed position, with Black to move. But Black is in zugzwang—he has no moves!
White wins
###148
1 Kg5    Nf2
2 h4!    Ne4+
3 Kg6    Nxd6
IF BLACK refuses the Bishop, then after 3 ... Nf2 4 h5, Ng4 5 Kg5, the Knight must retreat, since the King can no longer move and still protect it.
4 h5    Nc4
5 h6    Ne5+
6 Kg7
White wins
###149
TO WIN this, Black's Bishop must be driven off either of the two diagonals leading to White's f8, the Queening square of the Pawn.
The first step:
1 Bc3        Ba3
2 Bg7        Bb4
3 Bf8
Moving in front of the Pawn, even though it blocks the Pawn for the time being, to force Black's Bishop to leave the diagonal.
3 ...        Bd2
The second step:
4 Bc5        Bh6
Ousted from one diagonal, the Bishop seizes another, also leading to f8*.
5 Bd4
Prevents Black from playing 5 ... Bg7.
5 ...        Kf5
Black moves his King, since his Bishop must stay where it is.
6 Bg7
Now the Bishop moves beside the Pawn, assuring its advance next move.
White wins
###150
IN ORDER that his Pawn advance to e7 (after which it is assured of Queening) White must drive the opposing Bishop off either of the two diagonals leading to that square.
1 Be7
As in the previous example, White's first step is to move his Bishop in front of the Pawn. This will evict Black's Bishop from one diagonal.
1  ...        Be3
2 Bf8        Bg5
Black of course seizes the other diagonal leading to e7.
3 Bg7        Kd6
4 Bf6
The second step: the Bishop moves beside the Pawn to dislodge Black from his occupancy of the critical diagonal.
4 ...         Bd2
The alternative, exchanging Bishops, is obviously hopeless.
5 e7
White wins
###151
TO WIN this, White must do two things: maneuver his King to f7, and dislodge Black's Bishop from the path of the Pawn.
1 Bg5    Bf8
If 1 ... Bb2 (or anywhere else on the long diagonal) 2 Bf6 wins instantly.
2 Kf 6 !    Be7+
Or 2 ... Kg4 3 Bd2, Bc5 (3 ... Kh5 4 Kf7, and Black is out of moves) 4 Bc3, Bf8 5 Kf7, Kf5 6 Bd2 and White wins.
3 Kf7    Bf8
4 Be3
Black is in zugzwang—he must move, and he has no moves!
White wins
###152
1 Kh7!
TO PREVENT 1 ... Kg8, after which Black's King could not be evicted.
1 ...    Bb2
Or anywhere else on the long diagonal.
2 Bf4        Bd4
3 Bh6+        Ke8
4 Bg7
Moving in front of the Pawn, the Bishop drives the adverse Bishop off the long diagonal.
4 ...        Bc5
The only move, since 4 ... Be3 is refuted by 5 Bb2 followed by 6 g7, winning.
5 Be5        Bf8
In accordance with principle, White can now deploy his Bishop to f4, and then h6 alongside the Pawn, but there is a faster win available.
6 Bd6!
White wins. After 6 ... Bxd6 7  g7 is decisive.
###153
1 Bg7            Bg5
IF BLACK tries 1 ... Bd2, then 2 Bh6, Bc3 3 Bg5, Bg7 (to prevent 4 h6) 4 Be7!, Bc3 5 h6, Bd4 6 Bf6, and White wins.
2 Bh6
Moving in front of the Pawn to dislodge Black’s Bishop.
2 ...         Bf6

If 2 ... Be7 3 Be3, Bf8 4 Bd4, Kh4 (on 4 ... Kf4 5 Bg7 is conclusive) 5 Be5, Kg4 6 Bf6, Kf4 7 Bg7, and the Pawn has a clear road ahead.
3 Be3
Threatening to advance the Pawn.
3 ...          Bg7
4 Bg5    Bf8
5 Bf6    Kf4
6 Bg7
White wins
###154
1 Bh4        Kb6
ON A Bishop move instead, say
1... Be5, White continues with
2 Bf6, Bh2 3 Bd4, Bg3 4 Ba7, Bf4 5 Bb8, Be3 6 Bh2, Ba7 7 Bg1, and wins.
2 Bf2+        Ka6
To prevent White from moving his Bishop to a7 and b8, after which it is easy to dislodge the enemy Bishop. For example, if 2... Kc6 3 Ba7, Bg3 4 Bb8, Bf2 5 Bh2, Ba7 6 Bg1, and it is done.
3 Bc5!
Forces the opposing Bishop out of the corner, since his King must stay where it is.
3 ...         Bg3
4 Be7        Kb6
Otherwise, the maneuver 5 Bd8 and 6 Bc7 is decisive.
5 Bd8+        Kc6
Stops 6 Bc7, but the exposed position of Black's Bishop enables White to gain a tempo.
6 Bh4!
Had Black played 3 ... Bf4 earlier, then White's sixth move would have been 6 Bg5, or if 3 ... Be5, then White's sixth would have been 6 Bf6.
6 ...         Bh2
7 Bf2        Kb5
8 Ba7        Ka6
9 Bb8        Bg1
10 Bf4        Ba7
11 Be3
White wins
###155
1 Bc6!
IT IS important to blockade the Pawn. If instead 1 Ke5, c6 2 Bxc6, Kc7 3 Bg2, Na8! 4 Bxa8 (other moves do not affect Black's reply) Kb6     5 Bg2, Kxa7 and Black has a drawn.
1 ...         Kd8
2 Kf5!
The natural move 2 Kf7 (to force 2 ... Kc8 3 Ke8, and an easy win) is beautifully demolished by the subtle 2 ... Nc8 (attacking the Pawn) resulting in 3 a8=Q, and stalemate!
Another pitfall is 2 Ke5, when 2 ... Ke7 gives Black the opposition and a draw.
2 ...        Ke7
3 Ke5!
White seizes the opposition, the objective being to work his King around to b7.
3 ...        Kf7
4 Kd4        Ke6
5 Kc5        Ke7
It is clear that throughout all this, the Knight must stay rooted to the spot.
6 Bf3        Kd7
7 Bg4+        Ke7
8 Kc6        Kd8
9 Kb7
Black's position now comes apart. He must not move the Knight, dares not move the Pawn, and should not move his King.
White wins.
###156
THE win seems simple and straightforward—but it is easy to fall headlong into a hidden trap prepared by Black.
1 g5+        Ke7
2 g6        Ke8
At this point 99 out of 100 players would move 3 g7, confidently and unhesitatingly, only to be set back on their heels by 3 ... Bf8 (pinning the Pawn) followed by 4 ... Bxg6 (sacrificing the Bishop for the precious pawn), the resulting position being a draw.
3 Bd7+!    Kxd7
White meets other replies with
4 g7.
White wins
###157
WHITE must not be hasty and move 1 h7, as the reply 1 ... e4 ruins his winning chances.
1 Ba7!        Ba1
Naturally, if 1 ... Bxa7 2 h7 wins for White.
2 Kb1        Bc3
3 Kc2        Ba1
4 Bd4!
A brilliant sacrifice. If Black captures by 4 ... exd4, then 5 Kd3 blocks the Bishop's path and assures the Rook Pawn of Queening.
4 ...        Bxd4
5 Kd3        Ba1
On 5 ... Kg5 6 h7, e4+ 7 Kxd4, and the Pawn cannot be over-taken.
After the actual move, White must not play 6 h7, as after 6 ... e4+ the liberated Bishop will prevent White's Pawn from Queening.
6 Ke4
White wins, with Black Pawn blocked out, his Pawn cannot be headed off.
###158
THE King is urged to the top of the board, into a crowded position that can only end in mate.
1 Bf2+        Kh5
2 g4+        Kh6
3 Kf6        Kh7
Evidently not relishing 3 ... Bh7 4 Be3#
4 g5        Kh8
5 Bd4        Kh7
Or 5 ... Bh7 6 Kxf7#.
6 Ba1        Kh8
7 g6!        fxg6
If 7 ... Bh7 8 Kxf7#
8 Kxg6#
###159
ONLY one piece can stop White's passed Pawn—Black's Bishop. So the Bishop must be persuaded to leave the critical diagonals!

1 e5        Bb7
2 e6        Bc8
3 e7        Bd7
4 Bh3!
Without a moment's delay! If 4 Kf8 instead, 4... c5 wins for Black!
4 ...        Be8
5 Kf8        Bh5
6 Bg4
But not 6 Be6 (intending 7 Bf7) as Black's reply is 6 ... f3 giving him a draw.
6 ...        Bg6
Still keeping an eye on the Pawn.
7 Bf5+        Bxf5
8 e8=Q
White wins
###160
VARIOUS offers are made to decoy Black's rampant Bishop, and one of them must be accepted.
1 c6    Ba4
If 1 ... Bd5 2 Bc4 (pins the Bishop) Bxc4 (or 2 ... g3 3 Bxd5+ and the Knight Pawn is harmless) 3 c7 and White wins.
2 c7    Bd7
3 Ke7    Bc8
4 Kd8    Bf5
On 4 ... Bb7 5 Bg2, Bxg2 6 c8=Q wins.
5 Bd3    Be6
6 Bc4    Kf7
7 c8=Q
White wins with a fantastic pin at e6.
###161
BEAUTIFUL ending with minimal means.
1 d7        Rh8
If instead 1 ... Rg2+ 2 Kf7 (White's moves are intended to keep the Rook from reaching the first rank) Rf2+ 3 Ke7, Re2+ 4 Kd6, and the Rook is out of checks.
2 Kg7        Rb8
The only square available on the first rank.
3 Bc7 (Domination)
White wins
###162
1 b7        Rd3+
THE only way the Rook can get back to the first rank. If Black tries 1 ... Rf5+, then 2 Be5, Rf8 3 Bd6+ wins the Rook.
2 Ke6!    Rd8
3 Bc7    Rh8
4 Be5    Rd8
If Black moves 4 ... Re8+ instead, then 5 Kf7, Rd8 6 Bc7, Rh8 7 Bd6+ leads by a transposition of moves into the actual play.
5 Ke7    Rg8 
6 Kf7    Rd8
7 Bc7    Rh8 
8 Bd6+    Ka5 
9 Bf8

Cuts off the Rook from the Pawn, but does Black have another resource? 
9 ...     Rh7+ 
10 Bg7 
White wins, as the Pawn will Queen.
###163
1 c7        Re2+
THE Rook tries to return to the first rank, to prevent the Pawn's Queening. Should the Rook attempt to get behind the Pawn, then this would occur: 1 ... Rg5+ 2 Kd4, Rg4+ 3 Kc3, Rg3+ 4 Kb2, Rg2+ 5 Ka3, and White wins.
2 Kf6        Re8
3 Ba4!        Rg8
4 Kf7        Rh8
5 Kg7        Ra8
The only square left, as 5 ... Rc8 allows 6 Bd7+ winning the Rook.
6 Bc6        Ra7
The Rook could not remain on the first rank. This looks good though, as the Pawn is pinned and cannot advance.
7 Bd7+!
Cleverly interposing the Bishop, so that the Pawn is freed.
7 ...        Kh4
8 c8=Q
White wins
###164
AN INTERESTING battle between Bishop and Rook.
1 a7        Rf5+
The Rook must get back to the first rank, to stop the Pawn. It cannot do so by 1 ... Rh8 as 2 Bf6+ wins the Rook.
2 Ke2     Re5+
If 2 ... Rf8 3 Bf6+, Kc5   4 Be7+ wins the Rook.
3 Kd2        Re8.
4 Bf2+
Now it's the Bishop's turn to check.
4 ...    Ke5
5 Bg3+    Kf5
6 Bb8
Shuts off the Rook and wins for White.
###165
THERE is a great deal of brilliant play in this innocent-looking miniature.
1 b7        Ra5+
2 Kd6!
On 2 Kxe6 for example, Black draws by 2... Ra6+ (but not by 2... Rb5 3 Bc6+, Kd8 4 Bxb5, Kc7 5 Ba6, and White wins) 3 Kd5, Rb6 followed by 4 Kd4,  Rxb7.
2 ...         Rb5!
A subtle defense involving the sacrifice of the Rook and an offer of the Bishop!
There was no hope in 2 ... Ra6+ 3 Bc6+, and the Pawn will advance, since Black must drop all business to get out of check.
3 Bc6+        Kd8
4 Bxb5        Bc8!
The idea being that 5 b8=Q or 5 b8=R stalemates Black.
5 b8=B!
A clever under-promotion. Two Bishops do not generally win against one Bishop, but Black is cramped and in a mating net.
5 ...        Bg4
If the Bishop moves to the other diagonal by 5... Bb7, then 6 Bc7+, Kc8 7 Bd7#.
6 Bc7+        Kc8
7 Ba6#
###166
BEFORE we see how White wins, let us look at some interesting tries:
1 g7, Be6+ 2 Kxe6, Kxg7, and Black draws against the Rook Pawn and a Bishop that does not control the Pawn's Queening square.
1 Kf6, Be8 2 h4, Bxg6 3 Bxg6, and Black draws by stalemate.
C]    1 Bg8, Be6+ 2 Kxe6, Kxg6, and again Black gets to the corner and draws.
The solution is:
1 Kg8!        Be6+
A pretty alternative is 1 ... Bf5 2 g7! Bxh7+ 3 Kh8, Kg6 (the Bishop must obviously stay where it is) 4 h4, Kh6 5 h5, and Black is in zugzwang. He must move, though every move loses.
2 Kh8        Bf5
Against 2 ... Bd5, White proceeds with 3 g7, Be6 4 Bg8, Bf5 5 Bd5, Bh7 6 Be4, Bxe4 7 g8=Q and wins.
3 g7!        Bxh7
4 h3!
The key move! If 4 h4 instead, 4 ... Kg6 5 h5+, Kh6, and it is White who is in zugzwang, and must give up both Pawns.
4...        Kg6
5 h4        Kh6
6 h5

White wins: Black is out of moves.
###167
1 g6
THE first step in the process of promoting one of the Pawns is to bring them both to the sixth rank.
1 ...        Bc2
2 f5        Bb3
3 Kg5        Bc2
4 f6        Bb3
5 Bb4+        Kg8
On 5... Ke8, White comes around by 6 Kh6, and 7 Kg7 followed by 8 f7+ and 9 f8=Q, Black being unable meanwhile to do anything but watch.

6 Kf4
To swing the King around to f7 where it will help the Bishop Pawn's advance.
6 ...        Bc4
On 6 ... Bc2 7 f7+ wins at once.

7 Ke5        Bb3
8 Kd6        Kf8
Here too, if 8 ... Bc2 9 f7+, Kg7 10 Ke7 is decisive.
9 Kd7+        Kg8
10 Ke7        Bc4
11 f7+
White wins
###168
1 Bc4+        Ke7
2 Ke4!
PREMATURE would be 2 f5, Bg7 3 Kf4, Bh8 and Black has an easy draw. White would have no winning chances after the anti-positional 4 e6, while 4 f6+ allows the Bishop to sacrifice itself for the two Pawns.
2 ...        Bg7
3 Kf5        Bh6
4 Kg4!        Bf8
Choice is fearfully limited. If 4 ... Bg7 5 Kg5 leads to the actual play, while 4 ... Ke8 5 f5, Ke7 6 f6+, Ke8 7 e6 wins easily.
5 Kg5        Bg7
6 Kg6
White moves his King to a dominating position before advancing his Pawns.
6 ...        Bf8
On 6 ... Kf8 7 Kh7 wins the Bishop, while 6 ... Bh8 allows 7 Kh7, cornering the Bishop literally and figuratively.
7 f5        Ke8
8 f6
Note that White keeps his Pawns as much as possible on squares opposite in color to those controlled by his Bishop.
8 ...         Bc5
9 Kg7        Bf8+
10 Kg8        Bc5
11 e6        Bb4
12 Bb5+        Kd8
13 Kf7        Bc5
14 e7+
White wins
###169
IN CONVOYING the Pawns up the board, care must be taken that the enemy Bishop does not sacrifice itself for both Pawns, thereby forcing a draw.
1 Ke2        Kg4
2 Bel
Not at once 2 f3+, Kg3, and time will be lost evicting the King.
2 ...        Bd6
3 f3+        Kf4
4 g3+        Kf5
5 g4+
The Pawns, you will notice, occupy whenever possible squares opposite in color to those controlled by the Bishop. In this way, the Pawns and the Bishop dominate as many squares as possible. Here, for example, the Pawns attack the white squares, while the Bishop attacks the black.
5 ...         Ke6
If 5 ... Kf4 6 Bd2+, Kg3 7 g5, Be5 8 g6, Bg7 (otherwise 9 Bh6 followed by 10 g7 wins) 9 Ke3, Kh4 (clearly if 9 ... Bh6+ 10 Ke4, Bg7
11 Bf4+, Kg2 12 Be5 is decisive) 10 Ke4, Kh5 11 Kf5, Bd4 12 Bg5, Bg7 13 Bf6, Bf8 14 g7 wins.

6 Kd3        Kd5
7 Bd2        Bc7
8 f4        Bb6
To prevent White's intended maneuver 9 Ke3, 10 Kf3, 11 g5 and 12 Kg4
9 Bc3        Bc5
10 g5

The Pawns now occupy black squares (the color of those controlled by the Bishop). They do so when they can move without hindrance to the white squares.

10 ...        Bb6
11 g6        Ke6
12 Ke4        Bd8
13 f5+        Ke7
14 Kd5        Kf8

If 14 ... Bc7 15 f6+, Kf8 16 Bb4+, Kg8 17 f7+, and White wins.

15 Ke6        Bg5
16 f6        Be3
17 Bb4+
White wins.
###170
1 Bf1        Bg4
2 h4        Bf5
3 Kf2        Bg4
4 Ke3        Be6
5 Kf4        Bd7
6 Bd3
CAREFULLY avoiding 6 g4 (to continue with 7 Be2 and 8 g5+) as Black would sacrifice his Bishop for the Knight Pawn, and then have a draw against a Rook Pawn and a Bishop of the wrong color (one that does not control the Pawn's Queening square).
6 ...        Bh3
7 Bf5        Bf1
8 g4        Be2
Threatening to force a draw by 9 ... Bxg4.
9 g5+        Kh5
If 9 ... Kg7, White continues with 10 Bg4 followed by 11 h5.
10 Kg3!
Avoiding the pitfall 10 g6, Kh6 11 Ke5, Bh5 12 Kf6, Bxg6 13 Bxg6, stalemate.
10 ...         Bd1
11 Be4        Bb3
12 Bf3+    Kg6
13 Kf4    Bf7
14 h5+    Kg7
15 Ke5    Bb3
16 Be4
Note how possession is taken of the white squares before playing h6+.
16 ...    Bf7
17 h6+
The Pawns of course occupy whenever possible squares opposite in color to those controlled by the Bishop.
17 ...    Kh8
18 Kf6    Bh5
19 Bd5
Not 19 g6 allowing 19 ... Bxg6, and a draw, nor 19 Bg6 (to drive off Black's Bishop) since it blocks the Knight Pawn.
19 ...    Kh7
20 Bf7    Be2
21 g6+
White wins.
###171
1 f4+        Kd6
2 f5        Ke5
3 d4+        Kf6
4 Kf4        Bb3
5 Bc6        Bc2
6 Bd7

PROTECTS the BP and threatens to advance the QP.

6...        Bb3
7 Ke4        Bc4

If Black plays 7... Bc2+, then 8 Kd5, Bxf5 9 Bxf5, Kxf5 10 Kc6 wins for White.

8 d5        Bb3
9 Be6        Bc4
10 Kd4        Be2

If 10 ... Ba2 11 Kc5, Bb3 12 d6 wins.

11 d6        Bb5
12 d7        Ke7
13 f6+        Kd8
14 f7        Ke7
15 f8=Q+    Kxf8
16 d8=Q+

White wins
###172
A WIN for White seems optimistic, his Pawns look so shaky.
1 Kg5    Ng4
On 1 ... Kxe7, White replies 2 Kxf5, and Black's Knight is stranded. The continuation could then be: 2 ... Kf7 3 e6+, Ke7 4 Ke5, Ke8 5 Kf6, Kf8 6 e7+, Ke8 7 Bb5#.
2 Kxf5
But not 2 Bxg4, fxg4 3 Kxg4,  Kxe7, as after 4 Kf5, Kf7 5 e6+, Ke7! (in King and Pawn against King positions the defender blockades the Pawn whenever possible) 6 Ke5, Ke8! (so that the Pawn, when it reaches the seventh rank, can do so only with check) 7 Kf6, Kf8 8 e7+ (delaying the Pawn's advance does not help, as White cannot force a win) Ke8 9 Ke6 and Black draws by stalemate.
2 ...    Nxe5
3 Ke6!    N  moves
4 Bishop mates accordingly, either Bb5# or Bh5#.
###173
1 Kh6        Bf7
READY to meet 2 Bxh7 with 2 ... Bxc4, after which Black's Bishop can sacrifice itself for the remaining Pawn: and draw.
2 Bd3        Be6
Obviously if 2 ... Kg8, White takes the Rook Pawn with check. while 2 ... Bg8 is refuted by 3 g5, Bf7 4 Bxh7, Bxc4 5 g6 followed by 6 g7#.
3 g5        Bg8
4 Bxh7!

This can bowl a fellow over!
4 ...        Bxh7
There is no relief in 4 ... Bxc4 5 g6 and mate next move.
5 g6        Bxg6
All that is left. If 5 ... Bg8 6 g7#, or if 5 ... Kg8 6 gxh7+, Kh8 7 Kg6, b5 8 c5, b4 9 c6, b3 10 c7 b2 11 c8=Q#.

6 Kxg6              Kg8
7 Kf                  Kf8
8 Ke6                Ke8
9 Kd6                Kd8
10 Kc6                  Kc8
11 Kxb6              Kb8
12 c5                   Kc8
13 Kc6!

White wins
###174
1 g7        e1= Q
2 g8=Q        Kb7
THE King must flee. If Black tries aggressive action, then after   2 ... Qa5 3 Kd7+, Kb7 4 Qc8#, or if 2 ... Qh4+ 3 Kc7+ and mate follows next move.
3 Qb3+        Kc6
Or 3 ... Ka8 4 Qa4+, Kb7 5 5 Qb5+, Ka8 6 Qa6+, Kb8 7 Bd6#.
4 Qb6+        Kd5
5 Qb5!
A quiet move, leading to a remarkable position. The threat is 6 Bf2, discovering check and attacking the Queen.
It is Black's turn to move, but despite all the moves at his command, he is helpless. The proof:
A] 5 ... Ke6    6 Qe8+ win-ning the Queen.
B] 5... Ke4     6 Qe8+ win-ning the Queen.
C] 5 ...    Qe6    6 Qb3+ win-ning the Queen.  
D] 5 ...    Qe4    6 Qb7+ win-ning the Queen.
E] 5 ...    Qh4+ 6 Be7+ win-ning the Queen.  
F] 5 ...    Qg3 6 Bf2+  win-ning the Queen.  
G] 5...    Qc3 6 Bb4+  win-ning the Queen.  
H] 5 ...    Qd2 6 Qd7+  win-ning the Queen.  
I] 5 ...    Qh1 6 Qb7+  win-ning the Queen.  
J] 5 ...    Qd1 6 Qd7+  win-ning the Queen.  
K] 5 ...    Qc1 6 Be3+ win-ning the Queen.  
L] 5 ...    Qa1 6 Bd4+ Kxd4 7. Qe5+  sweet skewer of the Queen.

White wins.
###175
WHITE has two pretty winning possibilities, based on his opponent's choice of defense.
1 a7
White resists the tempting pin of the Knight, as after 1 Bc3, Kxa6 2 Bxb4, Black has been stalemated.
1 ...        Nc6
If 1 ... Ka6 2 Kb8, Nc6+ 3 Kc7, Nxa7 4 b4, Nc8 5 Kxc8, Ka7 6 Kc7, Ka6 7 Kc6, Ka7 8 Kxb5 and wins.
2 Kb7        Nxa7
3 Bc3+
The hasty 3 Kxa7 allows 3 ... Kb4 followed by 4 ... Kxb3, and Black draws.
3 ...         b4
4 Bd4
Threatens to capture the Knight.
4 ...         Nb5
 Saves the Knight, but at a fearful cost.
5 Bxb6#
###176
1 Ba3
IF WHITE pushes the Pawn instead, then after 1 a5, Ne7 2 a6, Nd5 3 a7, Nb6, and the Pawn can go no further.
The actual move is intended to prevent the Knight from approaching the Pawn by way of e7.
1 ...        f5
Frees the square f6 for the use of the Knight.
2 d5!
White in turn takes away a square from the Knight! Black is now forced to occupy d5 with a Pawn.
2 ...        cxd5
3 a5        Nf6
4 a6        Ne8
On 4 ... Nd7 (to meet 5 a7 with 5 ... Nb6) 5 Bc5!, Nxc5 6 a7 wins for White. With the actual move (4 ... Ne8) Black expects to reply to 5 a7 with 5 ... Nc7.
5 Bd6!
When instead, he is offered the Bishop—an unwelcome gift!
5 ...     Nxd6
6 a7
White wins
###177
A ROOK is almost always helpless against two connected passed Pawns on the sixth rank. Knowing this bit of endgame lore makes it easy to find the winning combination, including the repeated offers of the Bishop.
1 Bh5!         Kg3
If 1 ... gxh5 instead, 2 h7, Rc8 3 g6, and the Pawns cannot be stopped. One of them will become a Queen.
2 Bxf6!        Kf4
Here if Black had played 2... Rxg6, the continuation 3 h7, Ra6+ 4 Kb4, Ra8 5 g6 should be convincing.
3 h7        Rc8
Or 3 ... Rc3+ 4 Kb4, Rh3 5 Bh5!, Rxh5 6 g6 and White wins.
4 Be8!
The sacrifice of the Bishop gains a tempo for White's next move.
4 ...        Rxe8
5 g6
White wins, his Pawns being too strong for the Rook to cope with.
###178
1 e8=Q        Rxe8
2 Bf 8        Re2+
3 Kh3!
IF 3 Kg3 (to answer 3 ... Re3+ with 4 Kf4) Re6 4 g8=Q, Rg6+ 5 Qxg6, hxg6 lets Black get away with a draw.
3 ...         Re3+
4 Kh4        Re4+
5 Kh5        Re5+
6 Kh6        Re1
Now if White plays 7 g8=Q, Rh1+ 8 Kg7, Rg1+ wins the Queen.
7 Bc5!        Re8
Naturally, 7... Rh1+ 8 Kg5 (threatening to Queen with check) leaves Black no play.
Black's actual move restrains the Pawn for the time being.
8 Kxh7
Intending 9 Bf8 and 10 g8=Q.
8 ...        Rd8
Now if 9 Bf8, Rd7 pins the Pawn and draws, the Rook sacrificing itself next move for the Pawn.
9 Be7!        Rc8
10 Bf8
Threatens to Queen the Pawn next move.
10 ...         Rc7
Pins the dangerous Pawn. Has Black squeezed out a draw?
11 Bd6!
Pins the pinner and wins! White will capture the paralyzed Rook and then Queen the Pawn.
###179
1 a7        Bd5
IF NOT for Black's Bishop, White's Pawn could Queen. White therefore entices the Bishop away, and then sets up an impenetrable barrier between Bishop and Pawn.
2 c4        Bb7
The Bishop stays on the diagonal, avoiding such palpable traps as 2 ... Bxc4 3 a8=Q+.
3 Bf3        Bxf3
4 d5
Shuts the Bishop out and wins.
###180
WHITE has a strong passed Pawn as compensation for his opponent's extra Pawn. It is the threat of Queening this Pawn that makes life difficult for Black.
1 b3+!        Kxb3
Black must take this Pawn, as after 1 ... Ka5 2 b4+, Ka4 3 c4, Be8 4 b5 shuts his Bishop off the long diagonal and lets the passed Pawn advance unhindered to the Queening square. The alternative capture by 1 ... Bxb3 leads to 2 c4, Bc2 3 a7, Be4+ 4 Bf3 and White wins.
2 c4        Be8
3 Bf3        e4
Chases White's Bishop off the long diagonal so that Black can occupy it with his own Bishop.
4 Bh5!        Bc6
5 Bf7
Ready to meet 5 ... Kb4 with 6 Bd5, or 5 ... Ka4 with 6 Be8, in either case rendering Black's Bishop impotent.
5 ...        Ba8
6 Bd5        c6
Black avoids the exchange of Bishops, which loses quickly, but the move he plays seems effective since White must leave the diagonal.
7 c5+! 
But he chooses not to!
7 ...        cxd5
8 cxb6

White wins, his next move being 9 b7.
###181
BLACK threatens to draw by bringing his King and Pawn to the seventh rank and forcing the Rook to sacrifice itself for the Pawn.
To win this, White must capture the Pawn outright.
1 Kc6
The King hastens towards the Pawn, as there is not a moment to lose.
1 ...        d3
2 Kc5        Ke3
3 Kc4        d2
4 Kc3        Ke2
5 Ra2
Pins the Pawn. White captures it next move and wins.
###182
THE Rook by itself cannot win the Pawn (for nothing that is). But White's King is far away. Can he get over in time to help the Rook?

1 Kb7        h5
2 Kc6        h4
3 Kd5        Kg4
4 Ke4        Kg3
If 4 ... h3 5 Rg8+, Kh4 6 Kf4 (threatens 7 Rh8#) Kh5 7 Rh8+, winning the Pawn.
5 Ke3        h3
There is no hope in 5... Kh3 6 Kf3, and Black must abandon the Pawn, nor in 5 ... Kg4 6 Kf2, h3 7 Rh7, and again the Pawn falls.
6 Rg8+        Kh2
Or 6 ... Kh4 7 Kf3, Kh5 8 Kg3, and White wins the Pawn.
7 Kf2        Kh1
8 Ra8
White can win the Pawn by 8 Kg3, h2 9 Rh8, Kg1 10 Rxh2. but the actual move is faster since it forces mate.
8 ...        Kh2
Or 8 ... h2 9 Ra1#.
9 Rh8        Kh1
10 Rxh3#
###183
1 Kb7        g5
2 Kc6        g4
3 Kd5        Kf4
4 Kd4        Kf3
BLACK would lose quickly after 4 ... f3 5 Rf8+, Kg4 6 Ke3, Kh3 (6 ... g2 7 Kf2 wins) 7 Kf3 (threatens mate) Kh2 8 Rh8+ and the Pawn is gone.
5 Kd3        g3
6 Rf8+        Kg2
If 6 ... Kg4 7 Ke2, Kh3 8 Kf1, Kh2 9 Rh8#.

7 Ke2        Kg1
If 7 ... Kh2 8 Rg8, Kg2 (or 8 ... g2 9 Kf2, and the Pawn is doomed) 9 Rg7, Kh2 10 Kf3, and White wins the Pawn.
8 Rg8        g2
9 Kf3        Kh1
Last chance! If White pounces on the Pawn without thinking, he stalemates Black.
10 Kf2’
White wins
###184
1 Rg1+
WHITE does not start with 1 Rh1 as 1 ...  h5 in reply saves the Pawn and advances its career.
1 ...        Kf5
The only square, since moving to the Rook file would block his Pawn.
2 Rh1       Kg6
Thus the King has been driven back.
3 Kb2        h5
4 Kc3        Kg5
5 Kd2        h4
6 Ke2        Kg4
7 Kf2        h3
8 Rh2        Kh4
9 Kf3
White captures the Pawn next move and wins.
###185
1 Rg8+       Kf3
2 Rh8       Kg4
FORCED, as the Pawn must be protected. Black's pieces are now back  to the position in the diagram, but White has gained time to place his Rook in the best possible position in Rook endings—behind the passed Pawn.
3 Kb2       h4
4 Kc2       h3
5 Kd2       Kg3
6 Ke2!       Kg2

If 6... h2 7 Kf1, and Black must give up the Pawn.

7 Rg8+        Kh1
Playing for a stalemate possibility. There is no fight in 7... Kh2    8 Kf2, Kh1 9 Rg1+, Kh2    10 Rg3, Kh1 11 Rxh3#.
8 Kf3        h2
9 Kg3!
A little device for lifting the stalemate.
9 ...        Kg1
10 Kh3+    Kh1
11 Rd8        Kg1
12 Rdl+
White wins
###186
WHILE the finishing touch here is similar to that in the previous example, there is a fine point in White's first move.
It is clear that White's King must hurry down to the Pawn, but the natural series of moves does not succeed in winning the Pawn: 1 Ke6, Kb3 2 Kd5, a3 3 Kd4 (or 3 Rb6+, Kc2 4 Kc4, a2 5 Ra6, Kb2, and White cannot win) a2 4 Kd3, Kb2 5 Rb6+, Kc1 6 Ra6, Kb2, and Black has a draw.
1 Rb6!
Shuts the King in! Unable to move past the Rook file, Black cannot prevent White's King from coming in to attack the Pawn. True, White must still guard against allowing a draw by stalemate.
1 ...                Ka2
2 Ke6              a3
3 Kd5              Ka1
4 Kc4              a2
5 Kb3!             Kb1
6 Ka3+            Ka1
7 Re6               Kb1
8 Rel+
White wins
###187
1 Rf8+        Kg6
ON 1 ... Ke4, White's King swings around to the far side of the Pawn to win: 2 Kf6, g4 (if 2 ... Kf4 3 Kg6+, Kg4 4 Rf5) 3 Kg5, g3 4 Kh4, g2 5 Rg8, Kf3 6 Kh3, and the Pawn falls next move.
2 Rf6+        Kh5
Against 2 ... Kg7, play proceeds: 3 Ke6, g4 4 Kf5, g3 5 Rg6+ followed by 6 Rxg3 and wins.
3 Ke6         g4
4 Kf5        g3
Or 4 ... Kh4 5 Rh6+, Kg3 6 Rg6 and White wins the Pawn.
    5 Rg6    Kh4
Holds on to the Pawn—at fearful expense.
6 Rh6#
###188
WHITE'S threats of mate on one side of the board, facilitate his attack on the Pawns on the other side.
1 Rc7+        Kg8
If 1 ... Kh6 2 Rc2, and Black must give up his Knight Pawn to prevent mate.
2 Rg7+        Kh8
If 2 ... Kf8 3 Rb7, and the mate threat wins the Knight Pawn.
3 Rb7        a3
4 Kg6
White mates in two moves. after
4 ...         b1=Q+
5 Rxb1        a2
6 Rb8#
###189
1 Rg6!
HOLDS both Pawns fast! In nearly all endings the Rook's best position is behind the passed Pawns—his own or the opponent's.
1 ...        Kd7
2 Rg4!        g2
On 2 ... Ke6 instead, 3 Rxf4, Ke5 4 Rg4 wins.
3 Rxg2        Ke6
4 Rg5!
Separates Black's King and Pawn. The King is confined to the first three ranks, while the Pawn must not take another step.
4 ...        Kf6
Black dares not move his Pawn. If for example, 4 ... f3 5 Rg3 attacks and wins the Pawn.
5 Rc5        Ke6
    6 Kb7        Kf6
On 6 ... Kd6? 7 Rf5 wins, while 6 ... f3 loses the Pawn by 7 Rc3, f2 8 Rf3.
7 Kb6        Ke6
8 Kb5        Kf6
9 Kc4        Ke6
10 Kd3        Kf6
11 Ke4
White wins
###190
THE Pawns are so far advanced that it is useless to attack them directly. Black would simply allow one Pawn to be captured while the other moved on to become a Queen.
To win this, White keeps his opponent busy with mating threats, while his own King gradually moves over to the Pawns.
1 Rd2+        Kb1
Black avoids 1 ... Ka1, the penalty being 2 Kb3 (with any Black move) followed by 3 Rd1#.
    2 Kc3    Kc1
Pawn moves are answered as follows:
A] 2 ... g2 3 Rd1+, Ka2 4 Rg1, h2 (4... Ka3 5 Ra1#) 5 Rxg2+, Ka3 6 Rxh2.
B] 2 ...  h2 3 Rd1+, Ka2 4 Rh1, g2 (4 ... Ka3 5 Ra1#) 5 Rxh2, and the remaining Pawn, being pinned, will fall next move.
    3 Ra2    Kd1
If 3... Kb1 instead, then 4 Re2, g2 (a King move is of course unthinkable) 5 Re1+, Ka2 6 Rg1, and White wins both Pawns.

4 Kd3        Kc1
On 4 ... Ke1 5 Ke3 (threatens mate) Kd1 (or 5... Kf1 6 Kf3 followed by capturing the Pawns) 6 Kf3, g2 (6 ... h2 7 Kxg3) 7 Kf2 wins.
5 Ke3        h2
If 5 ... g2 6 Kf2 is decisive, while 5 ... Kb1 loses by 6 Re2, g2 7 Re1+, Kb2 8 Kf3.
6 Ra1+        Kb2
7 Rh1        g2
If 7 ... Kc3 8 Kf3 wins the Pawns.
8 Rxh2
White wins
###191
1 Rh2      a3
2 Ke5              a2
3 Kd4              Kb1
4 Kc3              a1=N
IF 4 ... a1=Q+ 5 Kb3, and Black must give up his Queen to avoid mate.
5 Rb2+    Kc1
6 Ra2                Kb1
Forced, as the Knight is immediately captured if it emerges.
     7 Rxa6           Nc2
 8 Re6!     Na3

The alternatives are amusing:
A] 8 ... Na1 9 Re2, and to avoid mate, Black must move the Knight, and lose it.
B] 8 ... Kc1 9 Re2, and to avoid mate, Black must give up his Knight at once.
9 Re1+    Ka2
10 Re2+    Ka1
If 10 ... Kb1 11 Kb3 threatens mate and wins the Knight.
11 Kb3    Nb1
Other moves of the Knight allow its capture or instant mate.
12 Ra2#!
###192
A DELIGHTFUL ending, in which Black's Pawn turns out to be a liability.
1 Rg5+                  Kf8
Clearly, if 1 ... Kh7 2 Rh5+ wins the Bishop.
2 Rh5              Bc7

The only move, the alternatives being:
A] 2 ... Bg3 (or Bg1) 3 Rf5+, Ke8  (on 3 . .  Kg7 4 Rg5+ wins the Bishop) 4 Rg5. and Black must give up the Bishop, to avoid mate.
B] 2 ...  Bf4 3 Rf5+ winning the Bishop
C] 2... Bb8 3 Rh8+ winning the Bishop

3 Kd7        Bb6
The only stopping-place, 3... Bg3 being met by 4 Rf5+ (since only g7 or g8 are available for the Black King) followed by 5 Rg5+, again winning the Bishop.
4 Rb5        Ba7
5 Ra5        Bb6
6 Ra8+        Kf7
7 Kc6
White wins the Bishop, having no flight square, and the game.
###193
HOW does White proceed? There is nothing but a draw after 1 Rxa2, while 1 Kc3, a1=Q+ 2 Kb3, Qa5! scuttles the mate threat and wins for Black.
1 Re1+        Kb2
2 Ra1!
Rather unexpected!
2 ...         Kb3
Postpones the capture, but sooner or later Black will have to take the Rook.
3 Kd2        Kb2
4 Kd1        Kxa1
On 4 ... Kb3 5 Kc1, Ka3 6 Kc2, Kb4 7 Rxa2 wins.
5 Kc2        a5
6 b6        a4
Black hopes to draw by stalemate if his King is hemmed in, while if the King is freed, his Pawn will threaten to Queen.
7 b7        a3
8 Kc3!        Kb1
9 b8=Q+        Kc1
If 9 ... Ka1 10 Qd6, Kb1 11 Qd1#.
10 Qf4+        Kd1 (or Kb1)
11 Qf1#
###194
WHITE'S own King being in the way, his Rook is unable to stop the passed Pawns. White must find some other means to win the position—or even to save the game!
1 Kd6!
The key move. If instead 1 Kd4, d2, or if 1 Kf4, d2 2 Rd6, e2 wins for Black.
1 ...        d2
2 Kc7        d1=Q
If 2 ... b6 instead, 3 Rxb6, d1=Q - or anything else followed by 4 Ra6#.
3 Ra6+!        bxa6
4 b6+        Ka8
5 b7+        Ka7
6 b8=Q#
###195
1 Rg5        h3

IF 1 ... b4+ 2 Kc4, h3 (on 2 ... b3 3 Kc5 forces mate) 3 Rg4 leads into the main line of play.
2 Rg4+         b4+
3 Kc4                     h2

On 3 ... b3 4 Kc5 discovers check and mate.
4 Rg3                      g1=Q
Here if 4 ... b3 5 Rxb3, h1=Q 6 Ra3#.
5 Ra3+!         bxa3
6 b3#
###196
1 Rf7        Bb1
THE only refuge. If 1 ... Bg6 instead, 2 Rg7 wins a piece, or if      1 ... Be4 2 Rf4 pins the Bishop.
2 g5        Nh6
Stalemates the Knight, and threatens by 3 Ra7+, Kb3 (3 ... Kb4 4 Rb7+ wins the Bishop) 4 Rg7 White winning it.
2 ...        Ba2
The only chance, since King moves lose as follows:

A]    2 ...    Kb5  3 Rb7+
B]    2 ...    Kb4  3  Rb7+
C]    2 ...     Kb3  3  Rg7
D]    2 ...Ka5 3 Rg7, Ba2          4 Ra7+
E]    2 ... Ka3 3 Kc3 (threatens mate) Ka4 4 Ra7+, Kb5. 5 Rb7+
In all these variations White wins a piece.
3 Ra7+    Kb3
4 Kc1
White wins. Black, completely helpless, must either give away his Knight or desert his Bishop.
###197
1 e7        Re4
SHOULD Black play 1 ... Rb8 instead, to stop the Pawn, then 2 Ra6+, Kb4 3 Rb6+ (diverting the Rook by force) Rxb6 4 e8=Q wins for White.
After Black's actual move (1 ... Re4), he threatens mate as well as capture of the Pawn.

2 Rh3+        Kb4
3 Rh4!
A powerful pin by the unprotected Rook.
3...         Rxh4
4 e8=Q
White wins
###198
1 h6        Kf6
IF BLACK tries 1 ... Rf8, then 2 h7, Rh8 (the King cannot help, as after 2 ... Kf6 3 Rf1+, Kg7 4 Rxf8 wins) 3 Kg5, Ke6 4 Kg6, Ke7 5 Kg7, and Black will have to give up his Rook for the Pawn.
2 h7        Kg7
Apparently Black has a draw, as after 3 Kg4, Rf8 4 Kg5, Rh8, he captures the Pawn. But White has an effective little combination that finishes Black.
3 h8=Q+!     Kxh8
4 Kg4+
White wins the Rook and the game
###199
AN ATTEMPT to drive the Rook off the Queen file could only succeed if White's Pawn stood on the fifth rank. Should White try it now, this would be the result: 1 Ke4, Re8+ 2 Kf5, Rd8 3 Ke5, Re8+ 4 Kd6, Rd8+ 5 Kc5, Rc8+, and White makes no progress.
But a little artifice gets the Pawn up a step, and enables White to carry out his idea.
1 d5!        Ka5
Black may not play 1 ... Rxd5+ as the reply 2 Kc4, threatening mate and attacking the Rook, is killing.
2 Kd4        Ka6
3 Ke5
Now that the Pawn is on the fifth rank, White can attack with his King, drive the Rook away, and advance the Pawn.
3 ...         Re8+
4 Kf6        Rd8
Or 4 ... Rf8+ 5 Ke7, and White proceeds as in the main line of play.
5 Ke6        Re8+
6 Kd7        Rh8
7 d6        Rh7+
8 Kc6        Ka5
To avoid being mated.
9 d7        Rh8
10 Kc7        Rh7
11 Kc8
White wins
###200
1 Kb8        Rb2+
2 Kc8        Ra2
3 Rg6+

WHITE'S idea is to drive the King down, rank by rank.

3 ...        Kc5

The alternatives are:
A] 3 ... Kb5 4 Kb7 followed by Queening the Pawn.
B] 3 ... Kd5 4 Kb7, Rb2+ 5 Rb6, Ra2 6 Ra6 and again the Pawn will Queen.

4 Kb7        Rb2+
5 Kc7        Ra2
6 Rg5+

Taking care to avoid 6 Rc6+, Kd5 7 Kb7 (intending to play 8 Ra6 next move) as 7 ... Rxa7+ brusquely forces a draw.

6 ...         Kc4
7 Kb7         Rb2+
8 Kc6         Ra2
9 Rg4+         Kc3
10 Kb6        Rb2+

Otherwise, White will continue with 11 Ka5, Ra2 12 Ra4, Rb2, followed by 12 a8=Q, White wins.

11 Kc5        Ra2
12 Rg3+        Kc2
13 Rg2+        Kb3
14 Rxa2        Kxa2
15 a8=Q+
White wins
###201
1 Kf4
THREATENS 2 Rg8+ followed by 3 a8=Q.
1 ...              Kf2
Hides behind White's King to evade the check.
2 Ke4        Ke2
Black must keep on hiding, as should he try 2 ... Ra4+ for example, the continuation would be 3 Kd3, Ra3+ 4 Kc2, Ra2+ 5 Kb3, Ra1 6 Rf8+, and White wins.
3 Kd4        Kd2
4 Kc5        Kc3
Black stays in the shadow of White's King. If he plays 4 ... Rc1+ instead, then 5 Kb4, Rb1+ 6 Ka3, Ra1+ 7 Kb2, Ra3? 8 Rd8+ wins for White.
5 Rc8!        Rxa7
Otherwise the Pawn advances.
6 Kb6+
White wins the Rook and the game.
###202
1 Rc8        Kd6
IF 1 ... Kd7 2 Rb8, Rh1 3 Kb7, Rb1+ 4 Ka6, Ra1+ 5 Kb6, Rb1+ 6 Kc5 and White wins.
2Rb8        Rh1
3 Kb7        Rb1+
4 Kc8        Rc1+
5 Kd8        Rh1
Threatens mate. White does not reply to this by 6 Ke8 as 6 ... Rh8+ 7 Kf7, Rh7+ wins the Pawn.
6 Rb6+        Kc5
If 6 ... Ke5 7 Kc7, Rh7+ 8 Kb8, Rh8+ 9 Kb7, Rh7+ 10 Ka6, Rh8 11 Rb8, and White wins.
7 Rc6+!
The only way to win. If instead 7 Re6, Ra1 8 Re7, Kb6 draws. or if 7 Ra6, Rh8+ 8 Ke7, Rh7+ 9 Kf8 (moving to e6 or f6 costs a Rook) Rh8+ 10 Kg7, Ra8 11 Kf7, Kb5 12 Ra1, Kb6 and Black draws by taking the Pawn.
7 ...         Kb5
If 7 ... Kd5 (certainly not 7 ... Kxb5 8 a8=Q+) 8 Ra6., Rh7+ 9 Kc7, Rh7+ 10 Kb6, Rh6+ 11 Kb5 wins.
8 Rc8        Rh8+
9 Kc7        Rh7+
10 Kb8
White wins
###203
ROOK endings are full of tactical tricks, and the one shown here occurs frequently.
1 a7        Rb1+
2 Ka6        Ra1+
3 Kb6        Rb1+
4 Kc5        Ra1
Or 4 ... Rc1+ 5 Kb4, Rb1+ 6 Kc3, Rc1+ 7 Kb2, Rc7 8 Rh8 (threatens to Queen the Pawn) Rxa7 9 Rh7+ and White wins the Rook.
5 Rh8        Rxa7
Otherwise the Pawn advances to the magic square.
6 Rh7+
White wins the Rook and the game.
###204
ALTHOUGH his King is cut off from the Queen side, Black puts up a tough battle and is not easily subdued.
1 Kb5        Rb8+
2 Kc6        Rc8+
3 Kb7        Rc1
4 a6        Rb1+
On 4 ... Ra1 5 Rb3 shields the King from checks, after which the Pawn marches up without hindrance.
5 Kc6        Rc1+
6 Kb5        Rb1+
7 Ka4        Rb8
Black rejects 7... Ra1+ as then 8 Ra3, Rh1 9 a7, Rh8 10 Kb5, Ra8 11 Kb6 wins easily.
8 Ka5        Ra8
Here, if 8 ... Rb1 9 Ra3    is decisive.
9 Kb6        Rb8+
10 Kc7        Rb1
11 Ra3        Rc1+
12 Kb6        Rb1+
13 Ka5        Rb8
14 a7        Ra8
15 Kb6
White wins
###205
1 Kb4        Re7
2 Ra3
WHITE avoids the exchange of Rooks, as after 2 Rxe7+, Kxe7 3 Kb5, Kd7 4 Kb6, Kc8 5 a5 (or 5 Ka7, Kc7 and White can make no progress) Kb8, and Black has an easy draw.
2 ...        Ke8
3 a5        Kd8
4 a6        Ra7
5 Kb5        Kc8
Black's King had been cut off from the Queen side, but the offer to exchange Rooks enabled it to move over and help in the fight against the passed Pawn.
6 Rh3        Kb8
Two alternatives are:
A] 6 ... Rc7 7 Rh8+, Kd7 8 a7, Rxa7 9 Rh7+ winning the Rook.
B] 6 ... Rd7 7 Rh8+, Kc7 8 a7, and the Pawn will Queen.
7 Rh8+
But not the tempting 7 Kb6, Rb7+! 8 Kc5 (8 axb7 stalemates Black) Rb1, and White has been swindled out of a win.
7 ...        Kc7
8 Rg8     Kd6
Compulsory, since 8... Kd7 loses a Rook by 9 Rg7+.
9 Kb6        Rf7
10 a7
White wins
###206
1 Rh7+    Kg5
2 g7
INTENDS 3 Kh8 followed by Queening the Pawn.
2 ...        Ra8+
3 Kf7        Ra7+
4 Ke6        Ra6+
5 Kd5!
The point! If at once 5 Ke5, Rg6, and White is in zugswang: if he moves the Rook, 6 ... Rxg7 draws, while a move by his King al-lows the reply 6 ... Kf6 and the Pawn will fall.

5 ...        Rg6
The Rook gets behind the Pawn, further checks being useless, and   5 ... Ra8 (to head off the Pawn) ineffective after 6 Rh8.
6 Ke5!
The difference is now apparent.  It is Black who is short of moves.  His Rook must stay where it is, to prevent the Pawn from advancing, while his King has only one legal move.
6 ...        Kg4
7 Rhl        Kg5
Or 7 ... Rg5+ 8 Ke4, Kg3 (if 8 ... Rxg7 9 Rgl+ wins the Rook) 9 Rg1+, Kh4       10 Rxg5, Kxg5 11 g8=Q+, and White wins.
8 Rgl+    Kh6
9 Rxg6+    Kxg6
10 g8=Q+
White wins
###207
1 Kb2        Rb8
2 Rb3        Kf7
3 Ka3!
PLAYING 3 b7 instead would be premature, as after 3. .. Ke7 4 Ka3, Kd7 5 Ka4, Kc7, the Pawn is lost.
3 ...        Ke7
4 Ka4        Kd7
5 Kb5        Kc8
On 5 ... Rb7 instead, White has several ways of winning:

A] 6 Rc3, Kd8 (if 6 ... Rb8 7 Rd3+, Kc8 8 Kc6, Rb7 9 Rd8+, Kxd8 10 Kxb7, Kd7 11 Ka7) 7 Ka6, Rb8 8 Ka7, Rc8 9 Rxc8+.
B] 6 Ka6, Kc8 7 Rc3+, Kb8 8 Rh3, Ka8 (if 8 ... Kc8 9 Rh8+ wins the Rook) 9 Rh8+, Rb8 10 b7#.
6 Rc3+           Kb7
If 6 ... Kd7 7 Rd3+, Kc8 (7 ... Ke7 8 Kc6 with an easy win) 8 Kc6, Rb7 9 Rd8+, Kxd8 10 Kxb7, and the Pawn will Queen.
7 Rc7+            Ka8
8 Ra7#!
###208
1 Rd4!
THE Rook does double duty here: it guards the Pawn, frees the King for aggressive action, and it shuts out the opposing King from the Queen side where he might take part in the struggle.
1 ...        Ke6
If Black tries exchanging Rooks (so that his King can get to the Queen side) the play would go thus: 1 ... Rd8 2 Rxd8, Kxd8 3 Ka4!, Kc8 4 Ka5 (not 4 Kb5, Kb7: and Black, having the opposition, draws) Kc7 5 Ka6, Kb8 6 Kb6, and White wins.
2 Kc4!    Rc8+
If 2 ... Ke5 3 Rd5+, Ke6 4 b5, Rc8+ 5 Rc5, Kd7 (on 5 ... Rxc5+ 6 Kxc5, Kd7 7 Kb6, Kc8 8 Ka7 wins) 6 b6 (but not 6 Rxc8, Kxc8 7 Kc5, Kc7, and Black has a draw) Rxc5+  7 Kxc5, Kd8 8 Kd6!, Kc8 9 Kc6, and White wins.
3 Kb5
This and the following moves con-stitute a zigzag maneuver, customary in this type of position, and worth noting.
3 ...     Rb8+
4 Kc6     Rc8+
5 Kb7     Rf8
6 b5
White wins
###209
THE fact that Black's King is cut off from the Queen side helps White in what might seem to be a difficult task, his Pawn being so far away from the Queening square.
1 Kc3    Rc8+
2 Kd4    Rb8
3 Kc4    Rc8+
4 Kd5    Rb8
If 4... Rd8+ 5 Kc6, Rc8+ 6 Kb7, Rc3 7 Rb1, Ke6 8 b4, Kc6 9 b5, and White wins.
5 Rb1    Kf6
If 5 ... Rb5+ 6 Kc6, Rb4 7 Kc5, Rb8 8 b4, Ke6 9 Kc6, Rc8+ 10 Kb7, and White has no troubles.
6 b4                Ke7
7 Kc6!    Kd8
8 Rdl+    Kc8
9 Rh1    Kd8
10 Rh8+
White wins a Rook and the game.
###210
1 Kd3!
THE plausible 1 d6 does not win, the continuation being 1 ... Rh6 2 Rd2 (if 2 d7, Rc6+ 3 Kb4, Rd6, and Black wins the Pawn) Rh8 3 Kd4 (here if 3 d7, Rd8 4 Kc4, Kb6 and the Pawn will fall) Kb7 4 Rc2, Rh4+, and White cannot escape perpetual check or loss of the Pawn.
1 ...         Rg4
2 Ke3        Rh4
3 d6        Rh6
4 Rd2!
The advance 4 d7 would be premature as 4... Re6+ in reply followed by 5... Rd6 will remove the Pawn.

4 ...         Rh8
5 d7        Rd8
6 Ke4         Kb7
7 Ke5         Kc7
8 Ke6         Rh8

Moving the King instead loses im-mediately, e.g. 8 ... Kc6 9 Rc2+, Kb7 10 Ke7, and Black must give up his Rook for the Pawn.
9 Rc2+        Kb7
Blockading the Pawn is suicide: 9 ... Kd8 10 Rc8#.
10 Rh2!
A brilliant surprise move, and the only way to win.
10 ...          Rg8
Clearly, 10 ... Rxh2 allows White to Queen his Pawn and win.
11 Kf7          Ra8
12 Ke7          Kc7
13 d8=Q+          Rxd8
14 Rc2+          Kb7
15 Kxd8
White wins
###211
THE Lucena position is almost 500 years old, but it is still one of the most important positions to know in the conduct of Rook endings.
White, as we can see, cannot win by simply moving his King out of the way, and Queening his Pawn. For example, if 1 Ke7, Re2+ 2 Kd6, Rd2+ 3 Kc6, Rc2+ 4 Kb5 (moving towards the Rook to escape the checks) Rb2+ 5 Kc4, Rd2, and Black wins the Pawn and draws.
Despite the apparent difficulties, White has two methods of winning:
1 Ra1
Intending to evict Black's Rook from the Bishop file so that White's King may emerge without being annoyed by checks.
1 ...         Kf7
2 Ra8        Rc1
If 2 ... Ke6 3 Ke8, Rh2 4 Ra6+, Kf5 5 d8=Q and White wins.
3 Rc8        Rd1
4 Kc7        Rc1+
5 Kb6        Rb1+
6 Kc5
White zigzags down the board until Black runs out of checks, then Queens his Pawn and wins.
The second method, beginning with the position in the diagram, is to Queen the Pawn by building a protecting bridge, as follows:
1 Rf4!    Rc1
2 Ke7     Re1+
3 Kd6    Rd1+
4 Ke6    Re1+
On 4 ... Rd2 5 Rf5 followed by 6 Rd5, shutting off Black's Rook, insures the advance of the Pawn.
5 Kd5    Rd1+
6 Rd4
White wins
###212
THE King, who is in check, must find a means of evading the persistent. checks by the Rook.

1 Ke6        Ra6+
On 1 ... Ra8 2 Rg1+, Kh7 3 Kf7, threatening mate as well as Queening the Pawn, wins for White.
2 Ke5+!

But not 2 Kf5, Kf7 3 Re1, Ra1 4 Re2, Re1, 5 e8=Q+ and White Queens and wins.
2 ...         Ra5+
3 Kf6        Ra6+
4 Kg5        Ra5+
Getting behind the Pawn by 4 ... Re6 yields to 5 Rf8+, Kg7 6 e8=Q, and White wins.
5 Kg6        Ra6+
6 Rf6        Ra8
7 Rd6

White wins, his Rook's next stop being d8—with or without check.
###213
1 Kd8        Rb8+
BLACK can attack with his King instead, with this result: 1 ... Kf6 2 e7!, Rb8+ (if 2 ... Rxe7 3 Rf1+, Ke6 4 Re1+ costs Black his Rook) 3 Kc7, Re8 4 Kd6, Rb8 5 Rf1+, Kg7 6 Kc7, Ra8 7 Ra1!, Rh8 8 Kd7, Kf7 9 Rf1+, Kg7 10 e8=Q and White wins.
2 Kc7        Rb2
3 Re1        Rc2+
4 Kd7        Rd2+
5 Ke8        Ra2
6  e7        Ra8+
7 Kd7        Ra7+
8 Kc6
White wins
###214
THE winning moves, in this ending of Chéron's, are as brilliant as they are forceful.
1 Kd3!
In reply to this, Black is practically forced to move his Rook. Moving his King instead leads to one of these pretty possibilities:
A] 1 ... Kh5 2 e6!, Ra6 3 Ke4! (in order to assist the Pawn by continuing with 4 Ke5 and 5 Kf6) Rxe6+ 4 Kf5, and White's attack on the Rook combined with a threat of mate on the move forces Black to give up his Rook.
B] 1 ... Kh5 2 e6!, Kh6 3 e7, Ra8 4 Ke4, Re8 5 Ke5!, Rxe7+ 6 Kf6, and again Black must lose his Rook or be instantly mated.
C] 1... Kh5 2 e6!, Kh4 3 Rg8, Ra7 (otherwise the Pawn simply moves on to e7 and e8) 4 Ke4, Kh5 5 Kf5, Kh6 6 Kf6, Kh7 7 Rd8, Ra6 8 Kf7, and the rest plays itself.
1 ...     Rb4 
Black's best defense is a move by his Rook, and that along the fifth rank.  Had he played something like 1 ... Ra6 instead (to prevent the Pawn from advancing) the win for White would be elementary. The con- tinuation would be 2 Ke4 and 3 Kf5, and Black could offer no resistance, his King being confined to the Rook file.
2 e6!        Rb6
3 Re1
    Threatens to advance the Pawn, and forces ...
3 ...        Rb8
4 e7        Re8
5 Kd4    Kg7
6 Kd5    Kf7
7 Kd6    Kf6
Threatens 8 Rf1+, Kg7, 9 Kd7, Ra8 10 e8=Q, and Black will have to give up his Rook.
7 ...      Ra8

Or 7 ... Rh8 8 Kd7, Re8         9 Rf1+, and Black must abandon his Rook.
8 Rf1+    Kg7
Naturally, if 8 ... Ke8 instead, the penalty is 9 Rf8#.  Gorgeous!
If White is hasty now and plays  9 Kd7 (threatening to promote the Pawn) Black replies with a series of checks from which the King cannot escape without loss of the Pawn. The play would be: 9 Kd7, Ra7+    10 Kd6, Ra6+ 11 Kd5, Ra5+ 12 Kc6, R-a6+, and if White approaches the Rook by 13 Kb7 or Kb5, the response 13 ... Re6 wins the Pawn.
The winning move brilliantly eliminates the possibility of perpetual check as a defense.

9 Ra1!    Rb8
Obviously if 9 ... Rxa1 instead, 10 e8=Q wins for White.
10 Kc7    Re8
11 Kd7    Kf7 

No better is 11 ... Rb8 12 e8=Q.
12 Rf1+
White wins the Rook and the game.
###215
THERE are delightful finesses in the position, innocuous though it looks.
1 Kd3!        Rd8+
On a King move instead, White simply pushes his Pawn up a square.
2 Kc4        Rc8+
3 Kd5        Rd8+
4 Ke5        Re8+
Must White lose time now by retreating, in order to save his Pawn?
5 Kf6!
Not at all, since he threatens instant mate.
5 ...         Rf8+
On 5... Kh5 instead, White replies 6 Re1 followed by advancing the Pawn.
6 Ke7        Rf5
7 e4!        Re5+
8 Kf6

Once again rescuing the Pawn by the threat of mate.
8 ...        Rh5
9 e5        Rh2
10 Rf1        Kh7
11 Kf7
White wins.
###216
ORDINARY tactics do not suffice, so White gives it a touch of brilliancy.
1 Rd2+
The King must be dislodged if White's King is to emerge.
1 ...        Ke7
White could now move 2 Kc7 but after 2 ... Rc3+, the King would have to return.
2 Rd6!    Rc3
If 2 ... Kxd6 instead, 3 Kc8, Rc3+ 4 Kd8, Rh3 (threatens mate) 5 b8, and the Pawn’s promoting with check wins for White.
3 Rc6!    Rxc6
4 Ka7
White wins
###217
THIS seems to be a dead draw, as White's Pawns are doomed, but White turns the dead draw into a neat win.
1 Kc4    Rxa6
Forced, as otherwise the King protects the Pawn and then escorts it up the board to be Queened.
2 a4!
Again threatening to proceed by 3 Kb5 and to march the Pawn up. 
2 ...        Rxa4+
3 Kb3!
Attacks the Rook on one hand, and threatens mate on the other.
White wins
###218
WHITE is three Pawns ahead, but threatened with loss of two of them by 1 ... Rxh3+ 2 Kb4, Kb7 and 3... Kxc7, after the White Rook has scurried away Black can draw against the remaining Pawn.
The win requires care and finesse.
1 Rd8        Rxh3+
Or 1 ... Kxc7 2 Rd3 and White has an easy win.
2 Rd3!        Rxd3+
Black must accept the offer, as the alternative 2 ... Rh8 leaves no hope after 3 Rc3.
3 Kc2!
If 3 Kc4 instead, 3 ... Rd1, and the threat of 4... Rc1 (with or without check) draws.
3 ...        Rd6!
Clever defense! If White plays 4 c8=Q, then 4 ... Rc6+ 5 Qxc6+, Kxc6 6 Kb3, Kb5, and Black, having the opposition, draws.

4 c8=N+!         Kc6
5 Nxd6         Kxd6
6 Kb3!         Kc6
7 Ka4         Kb6
8 Kb4         Kc6
9 Ka5
White wins
###219
SOME pretty ideas are concealed in a modest setting.
1 Rh3+    Kg7
If 1 ... Kg5 2 Rg3+, Kf5 3 Rxg6 Kxg6 4 a7 and White wins.
2 Rg3!    Rxg3
If Black refuses to capture, White exchanges Rooks, and Queens the Pawn.
3 a7    Rg1
4 Kb2
Prevents the Rook from moving to a1.
4 ...    Rg2+
5 Kxb3    Rg3+
6 Kb4
Moving to the Bishop file allows 6 ... Ra3, while 6 Ka4 lets Black draw by 6... Rg1 followed by 7... Ra1, with or without check.
6 ...     Rg4+
7 Kb5    Rg5+
8 Kb6    Rg6+
9 Kb7
White wins
###220
1 Rg1
IN ACCORDANCE with principle, the Rook supports the passed Pawn by moving behind it.
1...         Ra8
On 1 ... Rg3+ 2 Kf2, Kh5   3 g7, Kh6 4 g8=Q wins.
2 g7        Kh5
3 Kf3        Rg8
4 Kxf4        Kh6
5 Kf5        Kh7
If 5... Rxg7 instead, 6 Rh1#.
6 Kf6        Ra8
7 Rhl+        Kg8
8 Rh8#
###221
WHITE can gain a Pawn by 1 Rxf5 or 1 exf5, but the reply 1 ... Ra7+ leads to a drawing position.
The point of the winning combination is not to gain a Pawn, but to give one up!
1 Rg1+        Kh7
2 e5!        Rxe5+
Otherwise the passed Pawn is irresistible. For example, if 2 ... Ra7+ 3 Kf6 followed by 4 e6 wins easily, Black's King being cut off from the scene of action.
3 Kf7
Threatens mate on the move.
3 ...    Kh6
4 Kf6
Threatens mate again, whilst attacking the Rook.
White wins—first the Rook, then the game.
###222
1 Kg5!        Kh3
MAKES room for the Pawn to come through. Black can try to activate this Rook instead, with these possibilities:
A] 1 ... Rh8 2 Ra3+, Kf2 3 Kxg4, Rc8 4 Rc3, Ke2 5 Kf4, Kd2 6 Rc5, Kd3 7 Ke5 and White wins.
B] 1 ... Rh1 2 Rc2, Kf3 3 Rc3+, Ke4 4 c7, Rh8 5 c8=Q, Rxc8 6 Rxc8, Pg3 7 Kh4, g2 8 Rg8, Kf3 9 Kh3, and White wins.

2 Rh2+!    Kxh2
3 Kxh4        g3
4 c7        g2
5 c8=Q        g1=Q
6 Qh3#!
###223
THE move that breaks the pin is as brilliant as any I've seen in a long time.
1 b7    Ra7
2 Re1+    Kd8
3 Re7!
Frees the Pawn and threatens to transform it into a Queen.
3 ...    Kxe7
4 b8=Q
White wins. Black cannot hold on to his Pawn, since 4 ... Ra4 loses the Rook after 5 Qc7. Since Black’s King can only retreat to either e6 or e8, followed by 6 Qc6+.
###224
1 Ke7!
WHITE avoids the hasty 1 Rb8, since it allows Black to draw by 1 ... Rd7+ followed by 2 ... Rxa7.
1 ...        Ra8
Not 1 ... Rh8, which loses immediately by 2 Rb8.
2 Kd7!        Rf8
Clearly, 2 ... Kxa7 is fatal after the reply 3 Kc7, while 2...   Rxa7+ leads to the main line of play.
3 Rf2!        Ra8
On 3 ... Rg8 (or 3 ... Rh8) 4 Kc7 (intending to continue by 5 Ra2+ and 6 a8 - Queening) Ra8 5 Ra2+, Kb5 6 Kb7, and White wins.
4 Kc7        Rxa7+
5 Kc6
Threatens mate on the move.
5 ...        Ka5
6 Ra2+
White wins the Rook and the game.
###225
THE win is accomplished by a maneuver which forces Black's King to retreat rank by rank, until he is in position for the decisive blow.

1 Kb8
Threatens to Queen the Pawn.
1 ...             Rb2+
2 Ka8    Rc2
The Rook must return, to keep an eye on the Pawn.
3 Rf6+    Ka5
4 Kb8

Renews the procedure, Black's King meanwhile having had to move down a rank.
4 ...         Rb2+
5 Ka7        Rc2
6 Rf5+        Ka4
7 Kb7        Rb2+
8 Ka6        Rc2
9 Rf4+        Ka3
10 Kb6        Rb2+
11 Ka5        Rc2
12 Rf3+        Kb2
13 Rxf2        Rxf2
Black, unfortunately, may not capture White's Pawn in return.
14 c8=Q
White wins
###226
1 Rf8
INTENDING to continue with 1 ... Ke5 2 Rf6+, Kh5 3 Kg7, Rg5+ 4 Kf7, and the Pawn moves on to become a Queen and win.
1 ...        Rg5
2 Rf6+        Kh5
3 Rf5!        Rxf5
4 Kg7        Rg5+
5 Kf7
White avoids 5 Kf6, when 5 ... Rg6+ 6 Kf7, Rh6 wins for Black.

5 ...        Rf5+
6 Ke7         Re5+
7 Kd7         Rd5+
8 Kc7         Rc5+
9 Kb7         Rb5+

The last check, as the King hides behind the Pawn. 
10 Ka7 
White wins
###227
1 b7
A FORCEFUL beginning. A prelim- inary check instead leads to this: 1 Rc2+, Kd5 2 b7, Rh2+ 3 Kd1 (or 3 Kd3, Rh3+ followed by 4 ... Rb3) Rh1+ 4 Kd2, Rb1, and Black draws easily, his Rook now being behind the Pawn.
1 ...         Rh2+
2 Ke3        Rh3+
Rushing back to the first rank in stead is useless: 2 ... Rh8 3 Rc2+, Kd5 4 Rc8, and White wins.
With the actual move, Black hopes for 3 Kf4, Rb3 4 Rc2+, Kd5, and he has a drawn position.

3 Ke4!        Rb3
4 Rc2+        Kb5
Such a move must be distressing, Black having to obstruct the action of his own Rook.
5 b8=Q+
White wins
###228
1 Rg1+
WHITE does not start with the tempting 1 Rxg7 as after 1 ... Rh6 in reply he can make no progress.
1 ...        Kh2
2 Rg2+        Kh1
On 2 ... Kh3 instead, 3 Kg1 followed by 4 Rh2+ wins for White.
3 Kg3!
Prevents 3 ... Rxh7 as 4 Rd2 with a dire mate threat would be the penalty.
3 ...         Rh6
What else is there:
A] 3 ... Rg5+ 4 Kf3, Rxg2 5 h8=Q+, and White wins.
B] 3 ... Rg5+ 4 Kf3, Rh5 5 Rxg7, Kh2 6 Kf4, Rh6 7 Kg5, Rh3 8 Kg6, Rh4 9 Kf7 followed by 10 Kf8, and White wins.
C] 3 ... Rg5+ 4 Kf3, Rf5+ 5 Kg4 (threatens to Queen the Pawn) Rf8 6 Kg3 (now the idea is 7 Rh2+, followed by Queening) Rh8 7 Re2, and it's all over.
4 h8=Q!    Rxh8
5 Ra2
White wins
###229
A STRANGE duel of Rooks, where both sacrifice themselves to help their Kings!
1 e7        Rc8
An immediate offer of the Rook instead, leads to this: 1... Re5 2 Kxe5, e2 3 Rf3+, Kd2 (otherwise 4 Re3, and the Rook gets behind the Pawn) 4 Rf2, Kd1 5 Rxe2, and White wins.
2 Rf8        Rc6+
Black has no time for 2... e2 as White captures the Rook with check, in response.
3 Kf5        Re6!
Brilliant, and Black's best chance to save the game.
4 Kxe6        e2
5 Rf3+        Kd4
Here too, if 5... Kd2 6 Rf2 pins the Pawn and wins.
6 Rf4+        Kd3
To prevent 7 Re4--or so he thinks!
7 Re4!
"A Roland for an Oliver,” as the medieval Charlemagne saga “tit for tat” so amply describes.
7 ...        Kxe4
8 e8=Q        e1=Q
The position on the board (almost symmetrical) is artistic. 
9 Kf6+
White wins the Queen and the game with a beau of a discovered check.
###230
1 g7        Rd8
2 Re6+        Kf4
IF 2 ... Kd4 (or anywhere else on the Queen file) 3 Rd6+, Rxd6 4 g8=Q wins.
3 Rf6+        Kg3
Prevents White from playing 4 Rf8 as 4 ... Rd1+ in reply would be fatal.
4 Rg6+        Kh3
Ready to circumvent 5 g8=Q with 5 ... Rxg8 6 Rxg8, and Black is stalemated.
5 Kg1        Rg8
Now Black's threat is 6 ... Rxg7 7 Rxg7, and he draws by stalemate.
6 Kf2        Kh2
7 Rg4!        Kh3
The alternative is 7... h3  8 Rg3, Kh1 (a Rook move allows the Pawn to Queen) 9 Rxh3#.
8 Kf3        Kh2
9 Rxh4+        Kg1
10 Rh7        Rd8
Or 10 ... Kf1 11 Rh1#.
11 Rh8
White wins. A beautifully timed study.
###231
A BLOCKADING action to make Black's Rook impotent, is the winning idea in this ending.
1 Kc3
Threatens mate on the move.
1 ...        Ka4
Or 1 ... Ka2 2 g7, Rg2 3 Rb2+ (luring Black's Rook away from the Pawn) Rxb2 4 g8=Q+, and White wins.
2 Rb4+        Ka5
3 Rg4!        fxg4
4 g7
White wins, since his Pawn cannot be stopped.
###232
"THE Art of Exchanging Rooks,” by Vlk (improbable name!).
1 Kb8        Rc6
2 Rf6
Daring Black to exchange, the consequence of which would be 2 ... Rxf6 3 gxf6, c4 4 f7, c3 5 f8=Q, c2 6 Qc5, and White wins.
2 ...        c4
"Two can play at that game," as the fellow says. Now if White ex-changes Rooks, this happens: 3 Rxc6, dxc6 4 g6, c3 5 g7, c2 6 g8=Q, c1=Q, and the position is a draw.
3 Rh6+
Once more putting the question to Black. If now 3 ... Rxh6 4 gxh6,   c3 5 h7, c2 6 h8=Q, and the Pawn Queening with a check wins for White.
3 ...        Kg2
4 Rxc6
Now White exchanges, having forced Black's King to the Knight file.
4 ...        dxc6
5 g6        c3
6 g7        c2
7 g8=Q+
Queens with check, and White wins.
###233
A SUDDEN sacrifice makes the process of winning this ending a quick and easy one.
1 Rb7+        Ka3
Clearly if 1 ... Kc3 2 Rxc7+, followed by exchanging Rooks and then pushing the Pawn, wins on the spot.
2 g7        Rg1
3 Rxc7
Threatens to win by 4 Rc3+, Kb4 5 Rg6.
3 ...         Kb4!
The best defense. If instead 3 ... Kb2 4 Re7, (intending 5 Re2+ followed by 6 Rg2) forces 4 ... Kc1 when 5 Re1+, Rxe1 6 g8=Q wins for White.
4 Kh4        Rg5
Prevents White's King from moving closer to his Pawn, to help it Queen. An attempt to Queen his own Pawn leads to this: 4 ... f5 5 Kh5, f4 6 Kh6, f3 7 Rf7, Rg3 8 Kh7, Rh3+ 9 Kg8, Kc4 10 Kf8, and White wins.
5Re7
Intends 6 Re4+ followed by 7 Rg4.
5 ...        Kc5
There is no comfort in 5 ... Kc3 6 Re3+, Kd4 7 Rg3, and White wins.
6 Re5+!        Rxe5
What else is there? If 6 ... fxe5 7 Kxg5 wins, or if 6 ... Kd6 7 Rxg5, fxg5+ 8 Kxg5, and the game is over.
7 g8=Q
White wins
###234
WHITE is two Pawns behind, but his King and Rook are strongly placed. Black has two passed Pawns, but his King is confined to a corner, and faces great dangers.

1 Kf7
Threatens to win on the spot by  2 Rxh6 mate.
1 ...        Rxh5
If Black tries 1 ... Kh7 instead, then 2 Rg7+, Kh8 3 Kg6, f3 (other moves are no better) 4 Re7 forces mate.
2 Rg8+        Kh7
3 Rg7+        Kh8
4 Kg6        g4
To permit the Rook some freedom, Black must do something to counter White's threat of 5 Rd7 followed by mate.
5 Ra7        Rg5+
6 Kxh6        g3
What choice is there? Black may not play 6... Rg8, the penalty being 7 Rh7 mate. The Rook must stay on the Knight file, since moving it along the rank allows   7 Ra8+ and mate next.
7 Kxg5    g2
Or 7 ... Kg8 (to give his King room) 8 Kxf4, and the other Pawn will fall.
8 Ra1
Stops the Pawn from Queening, and threatens to win by the simple, brutal 9 Kxf4.
8 ...         f3
9 Kg6!
White wins, as mate will be forced.
###235
WHITE has only one Pawn to Black's four, but the strong position of his King suggests mating possibilities.
1 e6!
On the immediate 1 Kc7, Black defends by 1 ... a6, and after 2 Kb6, Rb5+ 3 Kxa6, Rb7, he (White) has nothing to fear.
1 ...    fxe6
But not 1... Ra6 (attempting to pin the Pawn) as 2 exf7, Rxh6 3 f8=Q+ wins.
2 Kc6    a6
If 2... Ra6+ instead, 3 Kc7 forces a quick mate.
3 Rh8+        Ka7
4 Rh7+        Ka8
The alternative leads to a pretty finish: 4 ... Kb8 5 Kb6 (threatens mate) Rb5+ 6 Kxa6, and the Rook is trapped.
5 Kb6        Rb5+
6 Kxa6        Rb8
The only square left unfortunately.
7 Ra7#
###236
"THE Skewer," illustrated here, is a tactical trick which appears quilt often in Rook endings.

1 g7        Rg6
2 Ra1
Intending 3 Ra8 and 4 g8=Q.
2 ...        Kb8
3 c7+
White offers a Pawn to draw the King away from the critical square.
3 ...        Kxc7
Other King moves lead to this:
3 ... Kc8 4 g8=Q+, Rxg8 5 Ra8+, and White wins the Rook.
3 ... Kb7 4 c8=Q+, Kxc8 5 g8=Q+, Rxg8 6 Ra8+, and White wins the Rook.
4 Ra8
One way or another, the Rook gets to this square. The threat now is Queening the Knight Pawn.

4 ...        Rxg7
Black's King and Rook are now in a straight line—in position to be skewered.
5 Ra7+
Wins the Rook and the game.
###237
WHITE is a Pawn ahead, but what to do? If 1 b7, Kc6, and he makes no headway, or if 1 Rb7, Kc6, and after the exchange of Pawns, Black has no trouble drawing.
There is a win though, beginning with a brilliant move:
1 Rd8+!        Kxd8
2 b7        Rb4!
The Pawn must be stopped, and Black can be brilliant too.
3 Kxb4        c5+
Now the position begins to look like a draw, as after 4 Kxc5, Kc7 5 Kb5, Kxb7, and Black has nothing to fear from the Rook Pawn.
4 Kb5!
White spurns the Pawn, and plays this, which is a killer.
4 ...        Kc7
5 Ka6
Threatens 6 Ka7 followed by Queening with check.

5 ...         Kb8
6 Kb6         c4
7 a4         c3
8 a5         c2
9 a6         c1=Q
10 a7#!
Who would have thought that the Rook Pawn (hiding behind the King in the diagrammed position) would be the one to strike the final blow?
###238
BLACK is subtly drawn into position to lose his Rook by the skewer attack.

1 d5        Ra8
2 d6        Kf7
3 d7        Ke7
There is nothing now in playing 4 d8=Q+, as after 4 ... Rxd8 5 Rxd8, Kxd8 6 a6, Kc7, and Black overtakes the Rook Pawn.
4 a6        Kd8
5 Rd6
Prevents Black's King from ernerg- ing, the consequence (after 5 ... Kc7) being 6 d8=Q+, Rxd8 7 Rxd8, Kxd8 8 a7, and White wins.
5 ...        Rb8
On 5 ... Ra7 instead, 6 Rxg6, Rxa6 7 Rg8+, Kc7 (or 7 ... Ke7 8 Rg7+) 8 d8=Q+ wins for White.
6 Kf3        Rb4
Purpose: to prevent White's King from advancing, and to get behind the Rook Pawn.
7 a7         Ra4
8 Rxg6
Clearing the way for 9 a8=Q+, Rxa8 10 Rg8+ winning the Rook.
8 ...        Kxd7
Or 8 ... Rxa7 9 Rg8+, Kxd7 (forced) 10 Rg7+, and the Rook falls.
9 Rg8
Threatens to Queen the Pawn next move.
9 ...     Rxa7
10 Rg7+

White wins with Black’s Rook skewered.
###239
1 Rh8
PINS the Pawn and threatens a quick win by 2 Rxh7+, Rxh7 3 gxh7.
1 ...         Rd2+
2 Kf1        Rd1+
If instead 2... Rg2 (getting behind the Pawn) 3 Rxh7+, Kg3 4 g7, Kf3 5 Rh3+, Rg3 6 Rxg3+, and White wins.
3 Ke2        Rg1
Keeps the dangerous Pawn under observation.
4 Rxh7+        Kg3
5 Rh1!        Rg2+
Clearly, Black cannot afford 5... Rxh1 6 g7, nor 5... Kg2 6 Rxg1+, in either case with a win for White.
6 Ke3        Kg4
Ready to meet 7 g7 with 7 ... Kf5, uncovering the Rook.
7 Rh2!        Rg3+
On 7... Rg1 instead, the winning play is 8 Kf2, Rg3 9 Rg2, forcing an exchange of Rooks.
8 Kf2        Rf3+
Other moves allow the decisive 9 Rg2.
9 Kg1        Rg3+
The alternative is 9... Re3 10 g7, Re8 11 g8=Q+, Rxg8 12 Rg2+, and Black's Rook comes off the board.
10 Rg2
Forces an exchange of Rooks, after which nothing can stop the Knight Pawn.
White wins
###240
THERE are several plausible ways to begin, but Black can refute them, viz:
A]     1 axb7, Rxb4 2 b8=Q, Rxb8, and White must take the draw        by 3 Rxf2.
B]     1 b5, Nd6 2 a7, Ra4 3 d6, Nc4+, and Black's next move will be 4... Ktxb6 with a draw as result.
The line of play that wins (and does so artistically) is this:
1 a7        f1=Q!
A necessary sacrifice if Black is to get behind the passed Pawn.
2 Rxf1        Rh2+
3 Kc1!        Ra2
The Rook Pawn looks like a goner.
4 b5!        Ra1+
Black does not take the Pawn immediately, as after 4 ... Rxa7 5 Rf6+, Kg5 6 Ra6, Rxa6 7 bxa6, White wins, the Knight being unable to prevent the Pawn from reaching the last square.
5 Kc2!
White is generous, as 5 ... Rxf1 is penalized at once by 6 a8=Q, and he wins.
5 ... Ra2+
6 Kb1   Rxa7
7 Rf6+  Kg5
The point is that the King must move to the fourth rank (which he does) or the second. Should he choose the latter, say by 7 ... Kg7, then the play would run: 8 Ra6, Nd6 (attacks the Pawn) 9 Rxa7+, and capturing with check saves the Pawn and wins for White.
8 Ra6 Nd6
Exchanging Rooks instead is hopeless, as 8 ... Rxa6 9 bxa6, Nd8 10 a7 and Black is a move too late to catch the Pawn.
9 Rxa7 Nxb5
10 Ra5
Pins the Knight (now we see why the King was forced to be fourth rank) and wins.
###241
1 f6!
AN ENERGETIC beginning. White threatens 2 fxg7+ winning a Rook.
1 ...     Rg8
Choice is limited: 1 ... Rxg6 instead allows 2 Ra8+ and mate next, while 1 ... gxf6 2 g7+ costs a Rook.
2 Rf7         d2
Black must depend on his passed Pawns, since the alternatives 2 ... gxf6 3 Rh7#, or 2 ... Rd8 3 Rxg7, d2 4 Rh7+, Kg8 5 f7+, Kf8 6 Rh8+, Ke7 7 Rxd8, Kxd8 8 f8=Q+, are not appetizing with quick Pawn elimination.
3 fxg7+        Rxg7
4 Kxh6        d1=Q
This seems to offer a glimmer of hope. Certainly there is none in 4 ... Rg8 5 Rh7#, nor in 4 ... Rxf7 5 gxf7, and Black's King must wait for the coup-de-grace.
5 Rf8+        Rg8
6 g7#
###242
FROM a game played in Moscow at the Alekhine Memorial, 7 Oct. 1956 between Botvinnik and Najdorf.  Screen shows position after the 63rd move.
By cleverly sacrificing a Pawn, White's King finds a good place to hide—right in the camp of the enemy!

1 e5        fxe5
2 fxe5
With this idea: 3 Rd7+, Rxd7 4 e6+, Ke7 5 exd7, Kxd7 6 Kg6, and White captures the abandoned Pawns and wins.
2 ...        Ke7
3 e6        Ra4
There is no fight in 3 ... Ra6 4 Rd7+, Kf8 5 Kg6!, Rxe6+ 6 Kh7, and Black's remaining Pawns will fall (6 ... Re4 7 Rxg7, Re6 8 Rg6).
4 g5!
A subtle sacrifice, the point of which we will see later. 
4 ...        hxg5
5 Rd7+        Kf8
6 Rf7+        Kg8
7 Kg6
Getting in between the opponent's Pawns, the King safeguards himself against troublesome checks.
7 ...        g4
8 h6!
Threatens 9 h7+, Kh8 10 Rf8#.
8...        gxh6
Or 8 ... Ra8 9 hxg7, g3 10 e7, Ra6+ (10 ... g2 11 Rf8+ and mate next) 11 Rf6 Rxf6+ 12 Kxf6, and White mates in two more moves.

9 e7        Ra8

On 9 ... Ra6+ 10 Rf6, Ra8 leads to the play that follows.

10 Rf6        Re8
11 Rd6
White wins, there being no specific against the effects of 12 Rd8, his next move.
###243
IN ENDINGS of Queen against a Pawn on the seventh rank, the method of winning shown here applies if the Pawn stands on the King, Queen, or Knight file. Should the Pawn occupy the Bishop or Rook file, different treatment is necessary, as we shall see in subsequent examples.

1 Qf2 

Pins the Pawn, and forces...
1 ...        Kd1
2 Qd4+        Kc2
3 Qe3        Kd1
4 Qd3+        Ke1
Black has been forced to block his Pawn, giving White's King time to move one square closer. White keeps repeating this maneuver until his King is near enough either to help capture the Pawn or to assist in bringing about mate.

5 Kf6        Kf2
6 Qd2        Kf1
7 Qf4+        Kg2
8 Qe3        Kf1
9 Qf3+        Ke1
10 Ke5        Kd2
11 Qf2        Kd1
12 Qd4+    Kc2
13 Qe3        Kd1
14 Qd3+    Ke1
15 Ke4        Kf2
16 Qf3+        Ke1
17 Kd3        Kd1
18 Qxe2+    Kc1
19 Qe2#
###244
THE customary procedure against a Pawn on the seventh rank does not work if the Pawn occupies the Bishop or the Rook file. For instance, if 1 Qd4, Kg2 2 Qg4+, Kh2 3 Qf3, Kg1 4 Qg3+, Kh1! and Black abandons the Pawn nonchalantly, since taking it allows a draw by stalemate.
White's winning chance in this and similar positions lies in his King being near enough (as it is here) to take part in a mating combination.

1 Kf4
Now Black's little trick of abandoning the Pawn does not work, as after 1 ... Kh1 2 Qe2 (for Heaven's sake, not 2 Kg3??, allowing 2 ... f1=N+, and Black wins the Queen) Kg2 3 Kg4, Kg1 (if 3 ... Kh1 4 Qf1+, Kh2 5 Qxf2+) 4 Kg3, f1=Q 5 Qh2#.
1 . .  .        f1=Q+
2 Kg3
The threat is 3 Qh2#. If Black's Queen moves (say to f4) to give the King room, then 3 Qg2#.
White wins
###245
1 Qb3
WHITE'S purpose is not only to pin the Pawn, but to prevent Black from reaching a1 with his King, in a try for stalemate.
1.  .  .    Kd2
2 Qb2    Kd1
3 Kf3!    Kd2
The response to Queening the Pawn (and it would come like a flash) would be 3 ... c1=Q 4 Qe2#.
4 Kf2    Kd1
Clearly, on 4 ... Kd3 5 Ke1 forces Black to give up his Pawn.
5 Qd4+        Kc1
6 Qb4        Kd1
7 Qe1#
###246
1 Qf5
PINS the Pawn, and cuts down the choice of reasonable reply to two moves by the King.
1 ...        Ka1
Against 1 ... Kb2, White proceeds 2 Qf2, Kb1 3 Kb3, c1=N+ (or 3... c1=Q, 4 Qa2#) 4 Ka3, Nd3 5 Qd2 and mate next move.
2 Kb3!    c1=Q
Or 2 ... c1=N+ 3 Ka3, Nd3 4 Qf1+ (side-stepping 4 Qxd3, stalemate) and mate next move.
3 Qa5+    Kb1
4 Qa2#
###247
1 Qb7+        Ka1
BLACK threatens to stalemate himself. White can win this type of position only if his King is near enough   to create a mating situation.
2 Qe4        Kb2
3 Qd4+        Kb1
Or 3 ... Kb3 4 Kd3, Ka3    5 Kc4, and Black will be mated next move.
4 Qd1+        Kb2
5 Qd2+        Kb1
Or 5 ... Ka1 6 Qcl#.
6 Kd1!        a1=Q
If 6 ... a1=N 7 Qb4+, Ka2 8 Kd2, Nb3+ 9 Kc3, Ka1 (a Knight move instead allows 9 Kc2, Ka1 10 Qa3#) 10 Kxb3 (definitely not 10 Qxb3) Kb1 11 Qe1#.
7 Qc2#
###248
WHITE wins because his King is close enough to assist in a mating operation. Were his King to stand on d6 instead, there could be no win.
1 Qh8+        Kb1
2 Qh1+        Kb2
3 Qh2+        Kb1
4 Kc4!
This move and the next gain two tempos for White, while Black is busy Queening his Pawn.
4 ...         a1=Q
5 Kb3
Threatens a mate which Black can only postpone by sacrificing his Queen.
White wins
###249
ORDINARILY this would be a draw, with White's King so far away, and the usual procedure against a Pawn on the seventh rank ineffective. For instance, if 1 Qe4, Kb2 2 Qb4+, Kc2 3 Qa3, Kb1 4 Qb3+, Ka1! Black has been forced to block his Pawn, but White cannot gain a move for his King to approach since he must release the stalemate.
White wins this though with a little device which enables his King to come two squares closer. Back to the diagrammed position!
1 Kb6!
Lifts the stalemate by screening the Queen.
1 ...         Kb2
2 Ka5+        Kc1
3 Qhl+        Kb2
4 Qg2+        Kb1
Or 4... Kb3 5 Qg7 (keeping Black King off the diagonal) followed by 6 Qa1, winning easily.
5 Ka4        a1=Q+
6 Kb3
And White mates (the position being an old friend by now, i.e. opposition with brutal force).
###250
1 Qg8+    Kf2
IF 1... Kh1 instead, 2 Qg3, a3 3 Qf2, a2 4 Qf1#.
2 Qh7        Kg3
3 Qd3+        Kg2
4 Qe4+        Kg3
If 4 ... Kf2 5 Qh1 wins, or if 4... Kg1 5 Qg4+, Kf2 6 Qh3, Kg1 7 Qg3+, Kh1 8 Qf2 and mate next move.
5 Kc5    a3
6 Kd4    a2
7 Qh1    a1=Q+
8 Qxa1    Kg2
9 Qb2+    Kg1

On 9 ...  Kg3 10 Qb7 followed by 11 Qh1 is decisive.

10 Ke3    h1=Q
11 Qf2#
###
`
