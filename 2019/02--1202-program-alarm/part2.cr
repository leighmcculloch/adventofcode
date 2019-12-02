program = [1,0,0,3,1,1,2,3,1,3,4,3,1,5,0,3,2,6,1,19,1,19,5,23,2,10,23,27,2,27,13,31,1,10,31,35,1,35,9,39,2,39,13,43,1,43,5,47,1,47,6,51,2,6,51,55,1,5,55,59,2,9,59,63,2,6,63,67,1,13,67,71,1,9,71,75,2,13,75,79,1,79,10,83,2,83,9,87,1,5,87,91,2,91,6,95,2,13,95,99,1,99,5,103,1,103,2,107,1,107,10,0,99,2,0,14,0]

def run_program(program : Array(Int32), noun : Int32, verb : Int32) : Int32
  i = 0
  program = program.dup
  program[1] = noun
  program[2] = verb
  while program[i] != 99
    case program[i]
    when 1
      program[program[i+3]] = program[program[i+1]]+program[program[i+2]]
    when 2
      program[program[i+3]] = program[program[i+1]]*program[program[i+2]]
    else
      raise "Encountered unknown opcode #{program[i]} at index #{i}"
    end
    i += 4
  end
  program[0]
end

def run_program_with_range(program : Array(Int32), noun : Range(Int32, Int32), verb : Range(Int32, Int32), want_output : Int32) : (NamedTuple(noun: Int32, verb: Int32) | Nil)
  noun.each do |n|
    verb.each do |v|
      output = run_program(program, n, v)
      return {noun: n, verb: v} if output == want_output
    end
  end
end

result = run_program_with_range(program, (0..99), (0..99), 19690720)
puts result

unless result.nil?
    puts "100 x noun + verb = #{100 * result[:noun] + result[:verb]}"
end
