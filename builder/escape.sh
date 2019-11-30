#!/usr/bin/env ruby

require 'cgi'
require 'yaml'

dir = 'output'

Dir.mkdir(dir) unless Dir.exist?(dir)

Dir['product_research/categories*'].each do |d|
  config = YAML.load_file(d)
  name = File.basename(d)

  config['marketplaces'].each do |m|
    m['categories'].each do |c|
      c['name'] = CGI.escape(c['name'])
    end
  end

  new_path = File.join(dir, name)
  File.open(new_path, 'w+') { |file| file.write(config.to_yaml) }
end
