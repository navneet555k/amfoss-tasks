  
require 'rubygems'                
require 'nokogiri'
require 'open-uri'

search = Nokogiri::HTML(open('https://www.google.com/search?q=linux&rlz=1C5CHFA_enIN871IN871&oq=linux&aqs=chrome..69i57j0l4j69i61l3.1704j0j7&sourceid=chrome&ie=UTF-8'))  
search.xpath('//a/div[@class="BNeawe vvjwJb AP7Wnd"]').each do  |top 10 search|
	puts ""
	puts top 10 search.text 
end

