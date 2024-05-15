FROM alpine:latest 

WORKDIR /usr/src/app 

COPY . /usr/src/app/

RUN chmod +x RioTradersPromoBot

CMD [ "./RioTradersPromoBot" ]