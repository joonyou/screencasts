# relative path #1
#-------------
#require './lib/foo'
#-------------
# relative path #2
#-------------
#require_relative 'lib/foo'
#-------------
# relative path #3
#-------------
#$LOAD_PATH << "."
#require 'lib/foo'
#-------------
# relative path #4
#-------------
#$:.unshift ".", "./lib"
#require 'foo'
#-------------
# absolute path
#-------------
#$:.unshift File.expand_path(File.dirname(__FILE__))
# - or -
#$:.unshift File.expand_path("..",__FILE__)
#require 'lib/foo'
#-------------

f = File.expand_path("..",__FILE__)
$:.unshift f, "#{f}/lib"
require 'foo'

puts Foo.hi
