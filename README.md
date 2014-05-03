# floodgate - reverse proxy that is able to suspend traffic without connection loss

NOTE: This is unusable for now, and it has never been battle tested.



Inspired by Braintree's Broxy architecture (http://drewolson.org/braintree_ha/presentation.html#slide-27).

Place floodgate between your nginx/thing and your application server.
Whenever you need to hold traffic temporarily, just use

    curl http://floodgate_url:port/hold

To open the floodgates again, use

    curl http://floodgate_url:port/release

You probably shouldn't be holding connections for more than a few seconds, but that's up to you.


## Missing features

- Configurability (remove the hardcoded paths and ports).
