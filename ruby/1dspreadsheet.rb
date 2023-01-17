# Auto-generated code below aims at helping you parse
# the standard input according to the problem statement.

n = gets.to_i
$map = n.times.map{gets.split(" ") }
$cache = {}
res = []

def find_val(arg)
    if arg =~ /\$(\d+)/
        ref = $1.to_i
        return $cache[ref] if $cache.has_key?(ref)
        res = f(ref, $map[ref])
        return res
        find_val(map[$1], map, cache)
    elsif arg =~ /\d+/
        return  arg.to_i
    end
end

def f(count, x)
    operation = x[0]
    arg1 = x[1]
    arg2 = x[2] if ! operation.eql?("VALUE")

    v_arg1 = find_val(arg1)
    v_arg2 = find_val(arg2) if ! operation.eql?("VALUE")

    case operation
    when "VALUE"
        res = v_arg1
    when "ADD"
        res = v_arg1.to_i + v_arg2.to_i
    when "SUB"
        res = v_arg1.to_i - v_arg2.to_i
    when "MULT"
        res = v_arg1.to_i * v_arg2.to_i
    end
    $cache[count] = res
    return res
end

count = 0
$map.each{
    |x|
    v = f(count, x)
    res.push(v)
    count += 1
}

res.each {|x| puts x}
