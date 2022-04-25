STDOUT.sync = true # DO NOT REMOVE
# Auto-generated code below aims at helping you parse
# the standard input according to the problem statement.


# game loop
loop do
    i = -1
    max = cpt = 0
    8.times do
        mountain_h = gets.to_i # represents the height of one mountain, from 9 to 0.
        if mountain_h > max
            max = mountain_h
            i = cpt
        end
        cpt += 1
    end
    
    # Write an action using puts
    # To debug: STDERR.puts "Debug messages..."
    STDERR.puts "test "
    puts "#{i}" # The number of the mountain to fire on.
end