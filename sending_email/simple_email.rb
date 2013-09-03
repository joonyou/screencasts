require 'time'

from = "donotreply@example.com"
from_name = "Joon"
to = "foo@example.com"
to_name = "Foo"
subject = "Evil Spam"
body = <<CONTENT
From: #{from_name} <#{from}>
To: #{to_name} <#{to}>
Subject: #{subject}
Date: #{Time.now.rfc2822}

Hi,

I hate hello world
CONTENT

server = "smtp.gmail.com"
port = 587
# username = from
password = "secret"

smtp = Net::SMTP.new(server,port)
smtp.enable_starttls
smtp.start(server, from, password, :login) do
  smtp.send_message(body, from, to)
end

