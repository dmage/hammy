#!/usr/bin/env ruby

require 'msgpack'

class EvalEnv
	def initialize(state, objectname)
		@state = state
		@objectname = objectname
		@cmdbuf = []
	end

	def get_state(key)
		elem = @state[key]
		if elem then
			elem["Value"]
		else
			nil
		end
	end

	def set_state(key, value)
		@state[key] = {
			"Value" => value,
			"LastUpdate" => Time.now.to_i
		}
	end

	def cmd(cmd, body)
		@cmdbuf << {
			"CmdType" => cmd,
			"Cmd" => body
		}
	end

	def objectname()
		@objectname
	end

	def key()
		@key
	end

	def data()
		@data
	end

	def results()
		{
			"CmdBuffer" => @cmdbuf,
			"State" => @state
		}
	end

	def eval(code, key, data)
		@key = key
		@data = data.sort { |x, y| x["Timestamp"] <=> y["Timestamp"] }
		return binding().eval(code)
	end
end

$stdin.binmode
u = MessagePack::Unpacker.new($stdin)
u.each do |obj|
	res = {}
	trigger = obj["Trigger"]
	if not trigger.empty? then
		e = EvalEnv.new obj["State"], obj["Key"]
		obj["IData"].each { |k, v|
			e.eval(trigger, k, v)
		}
		res = e.results()
	end

	res["CmdBuffer"] = nil if !res["CmdBuffer"] || res["CmdBuffer"].empty?
	res["State"] = nil if !res["State"] || res["State"].empty?

	$stdout.binmode
	$stdout.write(MessagePack.pack(res))
	$stdout.flush
end
