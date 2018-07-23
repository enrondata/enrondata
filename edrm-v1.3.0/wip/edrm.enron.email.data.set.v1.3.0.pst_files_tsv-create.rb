#!ruby

require 'multi_json'
require 'pp'

def add_sha(data = {})
  sha = IO.read('edrm.enron.email.data.set.v1.3.0.pst_files_shasum.txt')
  sha.split("\n").each_with_index do |line,i|
    next unless i > 0
    line.chomp!
    puts line
    parts = line.split
    if parts.length == 2
      if parts[0].length && parts[1].length
      	data[parts[1]] ||= {}
        data[parts[1]][:sha1] = parts[0]
        data[parts[1]][:name] = parts[1]
        #data[parts[1]][:url]  = 'https://s3.amazonaws.com/edrm.download.nuix.com/RevisedEDRMv1/' + parts[1]
      end
    end
  end
  data
end

def add_url(data = {})
  sha = IO.read('edrm.enron.email.data.set.v1.3.0.pst_files_urls.txt')
  sha.split("\n").each_with_index do |line,i|
    line.chomp!
    #puts line
    parts = line.split("\t")
    puts parts.length
    if parts.length == 2
      if parts[0].length && parts[1].length
        puts line
        puts parts[0]
        if parts[0] =~ /RevisedEDRMv1\/([^\/]+\.zip)$/
          filename = $1
          data[filename][:url] = parts[0]
          data[filename][:webindex_txt] = parts[1]
          if parts[1] =~ /^File\s+(\d+)\s*$/
            data[filename][:webindex] = $1.to_i
          end
        else
          raise "Warning"
      	end
      end
    end
  end
  data
end


data = {}

data = add_sha data
pp data
data = add_url data
pp data

def write_json(data)
  j = MultiJson.encode data, pretty: true
  File.open('enron.email.data.set.edrm-1.3.0.pst_files.json', 'w') do |fh|
    fh.puts j
  end
end

def write_tsv(data)
  byindex = {}
  data.each do |k,v|
    k2 = v[:webindex]
    byindex[k2] = v
  end
  #byindex = Hash[byindex.sort_by {|k,v| k.to_i }]
  File.open('enron.email.data.set.edrm-1.3.0.pst_files.tsv', 'w') do |fh|
  	fh.puts "#\tFile\tSHA1"
  	keys = byindex.keys
  	keys = keys.sort
    byindex.keys.sort.each do |k|
      v = byindex[k]
      fh.puts "#{k.to_s}\t#{v[:name]}\t#{v[:sha1]}"
    end
  end
end

write_json data
write_tsv data