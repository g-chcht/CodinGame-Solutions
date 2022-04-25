import sys
import math

def convert_game(game):
    frames = game.split()
    converted_frame = []
    for frame in frames:
        l = []
        for i in range(len(frame)):         
            if frame[i].isdigit():
                l.append(int(frame[i]))
            elif frame[i] == "/":
                l.append(10-l[-1])
            elif frame[i] == "X":
                l.append(10)
            elif frame[i] == "-":
                l.append(0)
            else:
                l.append("")
        converted_frame.append(l)

    return converted_frame

def compute_frame_score(j, frames):
    score = sum(frames[j])
    #strike
    if score == 10 and len(frames[j]) == 1:
        if len(frames[j+1]) > 1:
            score = score + frames[j+1][0] + frames[j+1][1]
        else:
            if (j < len(frames) - 2):
                score = score + frames[j+1][0] + frames[j+2][0]
            else:
                score = score + frames[j+1][0] + frames[j+1][1]
    #spare
    elif score == 10 and len(frames[j]) == 2:
        score += frames[j+1][0]
    return score

n = int(input())
for i in range(n):
    game = input()
    frames = convert_game(game)
    current_score, scores = 0, []

    for j in range(len(frames)):
        current_score += compute_frame_score(j, frames)
        scores.append(str(current_score))

    print(" ".join(scores))
