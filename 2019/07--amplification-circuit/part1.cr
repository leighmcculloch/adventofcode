class IntCodeComputer
  getter name : String
  getter phase_setting : Int64

  def initialize(@name, @phase_setting); end

  def run(program : Array(Int64), input : Int64) : Int64
    program = program.dup
    inputs = [@phase_setting, input]
    program_begin = nil
    commands = Array(String).new
    i = 0
    while program[i] != 99
      program_begin = program.dup if program_begin.nil?
      param_1_mode = program[i]//100 % 10
      param_2_mode = program[i]//1000 % 10
      param_3_mode = program[i]//10000 % 10
      opcode = program[i] - (program[i]//100*100)
      case opcode
      when 1
        v = value(program, program[i+1], param_1_mode) + value(program, program[i+2], param_2_mode)
        commands << program[i...i+4].to_s + " = #{v}"
        program[program[i+3]] = v
        i += 4
      when 2
        v = value(program, program[i+1], param_1_mode) * value(program, program[i+2], param_2_mode)
        commands << program[i...i+4].to_s + " = #{v}"
        program[program[i+3]] = v
        i += 4
      when 3
        program[program[i+1]] = inputs.shift
        commands << program[i...i+2].to_s + " = #{input}"
        i += 2
      when 4
        commands << program[i...i+2].to_s
        v = value(program, program[i+1], param_1_mode)
        #if v != 0
          #puts program_begin
          #commands.each { |c| puts c }
        #end
        return v
        commands.clear
        program_begin = nil
        i += 2
      when 5
        commands << program[i...i+3].to_s
        if value(program, program[i+1], param_1_mode) != 0
          i = value(program, program[i+2], param_2_mode)
        else
          i += 3
        end
      when 6
        commands << program[i...i+3].to_s
        if value(program, program[i+1], param_1_mode) == 0
          i = value(program, program[i+2], param_2_mode)
        else
          i += 3
        end
      when 7
        v = nil
        if value(program, program[i+1], param_1_mode) < value(program, program[i+2], param_2_mode)
          v = 1_i64
        else
          v = 0_i64
        end
        commands << program[i...i+4].to_s + " = #{v}"
        program[program[i+3]] = v
        i += 4
      when 8
        if value(program, program[i+1], param_1_mode) == value(program, program[i+2], param_2_mode)
          v = 1_i64
        else
          v = 0_i64
        end
        commands << program[i...i+4].to_s + " = #{v}"
        program[program[i+3]] = v
        i += 4
      else
        raise "Encountered unknown opcode #{program[i]} at index #{i} in #{@name}"
      end
    end
    -1_i64
  end

  def value(program : Array(Int64), param : Int64, mode : Int64) : Int64
    case mode
    when 0
      program[param]
    when 1
      param
    else
      raise "Encountered neither mode set: #{mode}"
    end
  end
end

def run_for_phase_setting(program, phase_settings, input)
  ics = [
    IntCodeComputer.new("A", phase_setting: phase_settings[0]),
    IntCodeComputer.new("B", phase_setting: phase_settings[1]),
    IntCodeComputer.new("C", phase_setting: phase_settings[2]),
    IntCodeComputer.new("D", phase_setting: phase_settings[3]),
    IntCodeComputer.new("E", phase_setting: phase_settings[4]),
  ]

  ics.each do |ic|
    input = ic.run(program, input)
  end

  input
end

program = File.read("input.txt").split(",").map { |n| n.to_i64 }
#program = [3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0].map{ |n| n.to_i64 }
input = 0_i64

max_output = nil
(0_i64..4_i64).to_a.each_permutation do |phase_settings|
  output = run_for_phase_setting(program, phase_settings, input)
  max_output = output if max_output.nil? || output > max_output
end
puts max_output
