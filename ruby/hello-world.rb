print "Hello World!\n"

empty_hash = Hash.new

default_hash = Hash.new('None')

puts default_hash[999]

hackerrank = {"Montreal" => 1032234, "London" => 7843003}

hackerrank.each do |key, value|
    puts value
end

hackerrank.each do |arr|
    puts arr[1]
end

# array manipulation 
# https://www.freecodecamp.org/news/common-array-methods-in-ruby/

def rot13(secret_messages)
    # your code here
      secret_messages.map {|x| x.tr("a-z", "n-za-m")}
  end
