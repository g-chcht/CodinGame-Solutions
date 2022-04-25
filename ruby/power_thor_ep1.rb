STDOUT.sync = true # DO NOT REMOVE
# Auto-generated code below aims at helping you parse
# the standard input according to the problem statement.
# ---
# Hint: You can use the debug stream to print initialTX and initialTY, if Thor seems not follow your orders.

# light_x: the X position of the light of power
# light_y: the Y position of the light of power
# initial_tx: Thor's starting X position
# initial_ty: Thor's starting Y position
@light_x, @light_y, @initial_tx, @initial_ty = gets.split(" ").collect {|x| x.to_i}

x = @initial_tx
y = @initial_ty

# game loop
loop do
    remaining_turns = gets.to_i # The remaining amount of turns Thor can move. Do not remove this line.
    
    # Write an action using puts
    # To debug: STDERR.puts "Debug messages..."
    if @light_x > x
        if @light_y > y
            x += 1
            y += 1
            puts "SE"
        elsif @light_y == y
             x += 1
            puts "E"
        else
            x =+ 1
            y =- 1
            puts "NE"
        end
    elsif @light_x == x
        if @light_y > y
            y += 1
            puts "S"
        else
            y -= 1
            puts "N"
        end
    else
        if @light_y > y
            x -= 1
            y += 1
            puts "SW"
        elsif @light_y == y
            x -= 1
            puts "W"
        else
            x -= 1
            y -= 1
            puts "NW"
        end
    end

    # A single line providing the move to be made: N NE E SE S SW W or NW
   # puts "SE"
end