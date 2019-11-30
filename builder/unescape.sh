#!/usr/bin/env ruby

require 'cgi'
require 'yaml'

Dir['../product_research/*'].each do |d|
  config = YAML.load_file(d)

  config['marketplaces'].each do |m|
    m['categories'].each do |c|
      c['name'] = CGI.unescape(c['name'])
    end
  end

  File.write(d, config.to_yaml)
end
