from __future__ import print_function, unicode_literals, division

from future.builtins import int
from future.builtins import range

import s2sphere as s2
import datetime
import json
from boto3 import Session
from requests_aws4auth import AWS4Auth
from elasticsearch import Elasticsearch, RequestsHttpConnection
from math import log
from random import randint


VERSION = '11-27-2018'
STAGING = 'production'
S3BUCKET = 'pinata-userfiles-mobilehub-1398053714'
ESHOST = 'https://search-pinata-es-wtz7zmneqgtg3g7bypilta3ske.us-east-1.es.amazonaws.com'
REGION = 'us-east-1'

class TemplateSuggestCollegeQuery(object):
    def __init__(self, esh, prefix, size=10, source=["unitid", "institution"]):
        self.__esh = esh
        self.__prefix = prefix
        self.__size = size
        self.__source = source

    def esh(self):
        self.__esh = esh
    def prefix(self):
        self.__prefix = prefix
    def size(self):
        self.__size = size
    def source(self):
        self.__source = source

    def query_suggest_college(self):
        return {'suggest': {'suggest': {'prefix': self.prefix(),
   'completion': {'field': 'combined'}}},
 'size': self.size(),
 '_source': self.source()}




class TemplateTopPostsQuery(object):
    def __init__(self, esh, location, seen_ids=[], filter_distance=50, filter_days=365*20):
        self.__esh = esh
        self.__location = location
        self.__seen_ids = seen_ids
        self.__filter_distance = filter_distance
        self.__filter_days = filter_days
        self.__size = size
        self.__from_doc = from_doc

    def esh(self):
        return self.__esh

    def location(self):
        return self.__location

    def seen_ids(self):
        return self.__seen_ids

    def filter_distance(self):
        return self.__filter_distance

    def filter_days(self):
        return self.__filter_days

    def size(self):
        return self.__size

    def from_doc(self):
        return self.__from_doc
    
    def query_top_posts(self):
        return {'query': {'bool': {'must_not': [{'ids': {'values': self.seen_ids()}}],
   'filter': {'geo_distance': {'distance': 50,
     'unit': 'km',
     'attachments.location': self.location()}}}},
 'sort': {'voteScore': {'order': 'desc'}},
 'size': self.size(),
 'from': self.from_doc()}

class TemplateExplorePostsQuery(object):
    def __init__(self, esh, location, seen_ids=[], filter_distance=50, filter_days=365, size=20, from_doc=0):
        self.__esh = esh
        self.__location = location
        self.__seen_ids = seen_ids
        self.__filter_distance = filter_distance
        self.__filter_days = filter_days
        self.__size = size
        self.__from_doc = from_doc

    def esh(self):
        return self.__esh

    def location(self):
        return self.__location

    def seen_ids(self):
        return self.__seen_ids

    def filter_distance(self):
        return self.__filter_distance

    def filter_days(self):
        return self.__filter_days

    def size(self):
        return self.__size

    def from_doc(self):
        return self.__from_doc

    def query_explore_posts(self):
        return {'query': {'function_score': {'query': {'bool': {'must_not': [{'geo_distance': {'distance': self.filter_distance(),
            'unit': 'km',
            'attachments.location': self.location()}},
          {'ids': {'values': self.seen_ids()}}],
         'filter': [{'term': {'active': True}},
          {'range': {'createdAt': {'gte': 'now-' + str(self.filter_days()) + 'd/d', 'lte': 'now/d'}}}]}}}},
     'sort': {'_script': {'type': 'number',
       'order': 'desc',
       'script': {'lang': 'painless',
        'params': {'absoluteCreationMillis': 0, 'redditConstant': 45000.0},
        'source': "(double)((((double)(Math.signum(doc['voteScore'].value))) * Math.log10(Math.max(Math.abs(doc['voteScore'].value),1))) + ((((int)(Math.round(((double)(doc['createdAt'].date.getMillis())) / 1000.0))) - params.absoluteCreationMillis) / params.redditConstant))"}}},
     'size': self.size(),
     'from': self.from_doc()}

class TemplateHotPostsQuery(object):
    def __init__(self, esh, location, seen_ids=[], filter_distance=50, filter_days=365, decay=0.99999, size=20, from_doc=0):
        self.__esh = esh
        self.__location = location
        self.__seen_ids = seen_ids
        self.__filter_distance = filter_distance
        self.__filter_days = filter_days
        self.__decay = decay
        self.__size = size
        self.__from_doc = from_doc

    def esh(self):
        return self.__esh

    def location(self):
        return self.__location

    def seen_ids(self):
        return self.__seen_ids

    def filter_distance(self):
        return self.__filter_distance

    def filter_days(self):
        return self.__filter_days

    def decay(self):
        return self.__decay

    def size(self):
        return self.__size

    def from_doc(self):
        return self.__from_doc

    def query_hot_posts(self):
        return {'query': {'function_score': {'query': {'bool': {'must_not': [{'ids': {'values': self.seen_ids()}}],
         'filter': [{'geo_distance': {'distance': self.filter_distance(),
            'unit': 'km',
            'attachments.location': self.location()}},
          {'term': {'active': True}},
          {'range': {'createdAt': {'gte': 'now-' + str(self.filter_days()) + 'd/d', 'lte': 'now/d'}}}]}},
       'functions': [{'script_score': {'script': {'lang': 'painless',
           'params': {'absoluteCreationMillis': 0, 'redditConstant': 45000.0},
           'source': "(double)((((double)(Math.signum(doc['voteScore'].value))) * Math.log10(Math.max(Math.abs(doc['voteScore'].value),1))) + ((((int)(Math.round(((double)(doc['createdAt'].date.getMillis())) / 1000.0))) - params.absoluteCreationMillis) / params.redditConstant))"}},
         'weight': 1},
        {'exp': {'attachments.location': {'origin': self.location(),
           'scale': '25km',
           'decay': self.decay()}},
         'weight': 1}],
       'boost_mode': 'replace',
       'score_mode': 'multiply'}},
     'size': self.size(),
     'from': self.from_doc()}

class TemplateNewPostsQuery(object):
    def __init__(self, esh, location, seen_ids=[], filter_distance=50, filter_days=365, decay=0.9999, time_weight=0.75, distance_weight=0.25, size=20, from_doc=0):
        self.__esh = esh
        self.__location = location
        self.__seen_ids = seen_ids
        self.__filter_distance = filter_distance
        self.__filter_days = filter_days
        self.__decay = decay
        self.__time_weight = time_weight
        self.__distance_weight = distance_weight
        self.__size = size
        self.__from_doc = from_doc

    def esh(self):
        return self.__esh

    def location(self):
        return self.__location

    def seen_ids(self):
        return self.__seen_ids

    def filter_distance(self):
        return self.__filter_distance

    def filter_days(self):
        return self.__filter_days

    def decay(self):
        return self.__decay

    def time_weight(self):
        return self.__time_weight

    def distance_weight(self):
        return self.__distance_weight

    def size(self):
        return self.__size

    def from_doc(self):
        return self.__from_doc

    def query_new_posts(self):
        return {'query': {'function_score': {'query': {'bool': {'must_not': [{'ids': {'values': self.seen_ids()}}],
     'filter': [{'geo_distance': {'distance': self.filter_distance(),
        'unit': 'km',
        'attachments.location': self.location()}},
      {'term': {'active': True}},
      {'range': {'createdAt': {'gte': 'now-' + str(self.filter_days()) + 'd/d', 'lte': 'now/d'}}}]}},
   'functions': [{'script_score': {'script': {'lang': 'painless',
       'source': "(doc['createdAt'].date.getMillis() / (double)((new Date().getTime()) +  86400000))"}},
     'weight': self.time_weight()},
    {'exp': {'attachments.location': {'origin': self.location(),
       'scale': '25km',
       'decay': self.decay()}},
     'weight': self.distance_weight()}],
   'boost_mode': 'replace',
   'score_mode': 'avg'}},
 'size': self.size(),
 'from': self.from_doc()}

class TemplateCollegePointQuery(object):

    def __init__(self, esh, location,  user_id=None, size=20, decay_type='linear', offset=0.0, scale=2.45, decay=0.5):
        self.__esh = esh
        self.__location = location
        self.__user_id = user_id
        self.__size = size
        self.__decay_type = decay_type
        self.__offset = offset
        self.__scale = scale
        self.__decay = decay

    def esh(self):
        return self.__esh

    def location(self):
        return self.__location

    def user_id(self):
        return self.__user_id

    def size(self):
        return self.__size

    def decay_type(self):
        return self.__decay_type

    def scale(self):
        return self.__scale

    def offset(self):
        return self.__offset

    def decay(self):
        return self.__decay

    def query_college_point(self):
        if self.user_id():
            geohash_aggregation_query = {'_source': 'location',
     'query': {'bool': {'must': {'match_all': {}},
       'filter': [{'term': {'byUserId': self.user_id()}}]}},
     'aggregations': {'geohash-aggregation': {'geohash_grid': {'field': 'location',
        'precision': 5,
        'size': 1}}}}
            try:
                geohash_hits = self.esh().es().search(index='events', doc_type='_doc', body=geohash_aggregation_query)
            except Exception as err:
                print(err)
                print("Could not retrieve user events")

    # def query_college_point(self):
    #     geohash_hits = []
    #     if self.user_id():
    #         geohash_aggregation_query = {'_source': 'location',
    #  'query': {'bool': {'must': {'match_all': {}},
    #    'filter': [{'term': {'byUserId': self.user_id()}}]}},
    #  'aggregations': {'geohash-aggregation': {'geohash_grid': {'field': 'location',
    #     'precision': 5,
    #     'size': self.size()-1}}}}
    #         try:
    #             geohash_hits = self.esh().es().search(index='events', doc_type='_doc', body=geohash_aggregation_query)
    #         except Exception as err:
    #             print(err)
    #             print("Could not retrieve user events")
            

    #     user_geofunctions = []
    #     location_weight = 1
    #     if geohash_hits and geohash_hits['hits']['total'] > 0:
    #         buckets_hits = geohash_hits['aggregations']['geohash-aggregation']['buckets']
            
    #         for bucket in buckets_hits:
    #             if location_weight < bucket['doc_count']:
    #                 location_weight = bucket['doc_count']
    #             user_geofunctions.append(
    #                 {
    #                 self.decay_type() : {
    #                   'location': {
    #                     'origin': bucket['key'],
    #                     'scale': str(self.scale()) + 'km',
    #                     'offset': str(self.offset()) + 'km',
    #                     'decay': self.decay()
    #                   }
    #                 },  
    #                 'weight': bucket['doc_count']
    #             })

    #     return {'query': {'bool': {'should': [{'function_score': {'query': {'bool': {'must': {'match_all': {}}}},
    #       'functions': user_geofunctions,
    #       'score_mode': 'max'}},
    #     {'function_score': {'query': {'bool': {'must': {'match_all': {}}}},
    #       'functions': [{'exp': {'location': {'origin': self.location(),
    #           'scale': '25km',
    #           'decay': 0.5}},
    #         'weight': location_weight}],
    #       'score_mode': 'max'}},
    #     {'bool': {'must': {'match_all': {}},
    #       'filter': {'geo_shape': {'ccwpath': {'shape': {'type': 'point',
    #           'coordinates': [self.location()['lon'], self.location()['lat']]},
    #          'relation': 'intersects'}}},
    #       'boost': location_weight*2+1}}]}}}


class TemplateUser(object):

    def __init__(self, esh, user_id, country_code, password):
        self.__esh = esh
        self.__password = password
        self.__user_id = user_id
        self.__country_code = country_code
        self.__access_code = 'user'
        self.__user_handle = self.access_code() + str(randint(0, 9999999))
        self.__unverified = True
        self.__ban_active = False
        self.__notification_default = True
        self.__notification_frequency = 'Weekly'

    def esh(self):
        return self.__esh

    def user_id(self):
        return self.__user_id

    def password(self):
        return self.__password

    def user_handle(self):
        return self.__user_handle

    def access_code(self):
        return self.__access_code

    def country_code(self):
        return self.__country_code

    def unverified(self):
        return self.__unverified

    def ban_active(self):
        return self.__ban_active

    def notification_default(self):
        return self.__notification_default

    def notification_frequency(self):
        return self.__notification_frequency


    def blank_user(self):
        return {
            'userId': self.user_id(),
            'marks': {
                 'active': self.esh().day_before(),
                'registration': self.esh().day_before(),
                'engagement': self.esh().day_before(),
                'vote': self.esh().day_before(),
                'share': self.esh().day_before(),
                'reply': self.esh().day_before(),
                 'post': self.esh().day_before(),
                 'report':self.esh().day_before(),
                 'referral':self.esh().day_before()
            },
            'properties': {
                'createdAt': self.esh().milliseconds(),
                'countryCode': self.country_code(),
                'username': self.user_handle(),
                'password': self.password(),
                'accessCode': self.access_code(),
                'unverified': self.unverified(),
                'modifiedAt': self.esh().milliseconds(),
                'banActive': self.ban_active(),
                'banEnd': 0,
                'accountActive': True
            },
            'attachments': {
                'counts': {
                    'banRecieved': 0,
                    'karma': 0,
                    'locationFollowing': 0,
                    'post': 0,
                    'reply': 0,
                    'reportRecieved': 0,
                    'reportTo': 0,
                    'share': 0,
                    'block': 0,
                    'friend': 0,
                    'trophy': 0,
                    'vote': 0
                },

                'profile': {
                    
                    'notifications': {
                        'notificationFrequency': self.notification_frequency(),
                        'friendPost': self.notification_default(),
                        'replyPost': self.notification_default(),
                        'trendingActivity': self.notification_default()
                    }
                    
                }
            }
        }

class TemplatePost(object):

    def __init__(self, esh, post_id, post, location, user_id):
        self.__esh = esh
        self.__post = post
        self.__post_id = post_id
        self.__user_id = user_id
        self.__location = location
        self.__post_type = self.post().get('type')
        self.__children = self.post().get('children')
        self.__attachments = self.post().get('attachments')
        self.__meta = self.post().get('meta')
        self.__access_code = 'review'
        self.__event_type = 'post'


        if self.post_type() == 'video':
            self.__event_code = 'postVideo'
            self.__event_media = 'video'
        elif self.post_type() == 'image':
            self.__event_code = 'postImage'
            self.__event_media = 'image'
        elif self.post_type() == 'text':
            self.__event_code = 'postText' 
            self.__event_media = 'text'

        self.__attachments['location'] = self.location()

        try:
            college_hits = self.esh().es().search(index='college-info', doc_type='_doc', body=TemplateCollegePointQuery(esh=self.esh(),location=self.location()).query_college_point())
            if college_hits.get('hits') and college_hits['hits']['total'] >= 1:
                self.__attachments['locationName'] = college_hits['hits']['hits'][0]['_source']['institution']
                self.__attachments['locationAlias'] = college_hits['hits']['hits'][0]['_source']['alias']
                self.__attachments['cityName'] = college_hits['hits']['hits'][0]['_source']['city']
                self.__attachments['unitid'] = college_hits['hits']['hits'][0]['_source']['unitid']
        except Exception as err:
            print(err)
            print('Could not query colleges for post')
        
        self.__level_30 = s2.CellId.from_lat_lng(s2.LatLng.from_degrees(self.location()['lat'], self.location()['lon']))

    def post(self):
        return self.__post

    def location(self):
        return self.__location

    def esh(self):
        return self.__esh

    def level(self, level):
        return self.__level_30.parent(level).to_token()

    def post_id(self):
        return self.__post_id

    def user_id(self):
        return self.__user_id

    def post_type(self):
        return self.__post_type

    def children(self):
        return self.__children

    def attachments(self):
        return self.__attachments

    def meta(self):
        return self.__meta

    def event_type(self):
        return self.__event_type

    def event_code(self):
        return self.__event_code

    def event_media(self):
        return self.__event_media

    def access_code(self):
        return self.__access_code

    def gen_create_post(self):
        return {'postId': self.post_id(),
                'createdAt': self.esh().milliseconds(),
                'modifiedAt': self.esh().milliseconds(),
                'eventCode': self.event_code(),
                'mediaType': self.event_media(),
                'eventType': self.event_type(),
                'userId': self.user_id(),
                'active': True,
                'accessCode': self.access_code(),
                'replyCount': 0,
                'shareCount': 0,
                'voteScore': 0,
                'downScore': 0,
                'upScore': 0,
                'replayCount': 0,
                'viewTime': 0.0,
                'attachments': self.attachments(),
                'meta': self.meta(),
                'children': self.children()
            }

    def gen_create_event_location(self):
        return {
                'eventCode': self.event_code(),
                'eventType': self.event_type(),
                'mediaType': self.event_media(),
                'createdAt': self.esh().milliseconds(),
                'modifiedAt': self.esh().milliseconds(),
                'aboutId': self.post_id(),
                'byUserId': self.user_id(),
                'location': self.location()      
            }


# class UtilitiesTracking(object):

#     def __init__(self, esh, user_id, user=None, location=None, registration=False, active=False, post=False, reply=False, referral=False, report=False, share=False, vote=False, engagement=False):
#         self.__esh = esh
#         self.__user_id = user_id
#         self.__user = user
#         self.__location = location
#         self.__parameters = {'registration': registration, 'active': active, 'post': post, 'reply': reply, 'referral': referral, 'report': report, 'share': share, 'vote': vote, 'engagement': engagement}

#     def esh(self):
#         return self.__esh

#     def parameters(self):
#         return self.__parameters

#     def user_id(self):
#         return self.__user_id

#     def user(self):
#         return self.__user

#     def location(self):
#         return self.__location


#     def update_tracking(self):
#         if not self.user():
#             try:
#                 user_hit = self.esh().es().get(index='users', id=self.user_id(), doc_type='_doc')
#             except Exception as err:
#                 raise AssertionError('Could not access users')

#             assert user_hit['found'], 'User Not Found'

#             self.__user = user_hit['_source']

#         trackingscriptline = ''
#         userscriptline = ''
#         for param in self.parameters():
#             if self.parameters()[param]:
#                 trackingscriptline += 'ctx._source.metrics.all.' + param + '++;'
#                 if self.user()['marks'][param] != self.esh().day():
#                     trackingscriptline += 'ctx._source.metrics.daily.' + param + '++;'
#                     userscriptline += 'ctx._source.marks.' + param + ' = params.date;'

#         if userscriptline:
#             script_update_user_mark_date = {
#                     'script': {
#                         'source': userscriptline,
#                         'lang': 'painless',
#                         'params': {
#                             'date': self.esh().day()
#                         }
#                     }
#                 }

#             try:
#                 self.esh().es().update(index='users', id=self.user()['userId'], doc_type='_doc', body=script_update_user_mark_date)
#             except Exception as err:
#                 print(err)
#                 print('Could not update user mark date')

#         if trackingscriptline:
#             script_update_registrations = {
#                     'script': {
#                         'source': trackingscriptline,
#                         'lang': 'painless'
#                     }
#                 }


#             create_tracking_element = {
#                 'date': self.esh().day(),
#                 'metrics': {
#                     'daily':{
#                         'registration': 0,
#                         'engagement': 0,
#                         'post': 0,
#                         'reply': 0,
#                         'active': 0,
#                         'referral': 0,
#                         'share': 0,
#                         'report': 0,
#                         'vote': 0
#                     },
#                     'all': {
#                         'registration': 0,
#                         'engagement': 0,
#                         'post': 0,
#                         'reply': 0,
#                         'active': 0,
#                         'referral': 0,
#                         'share': 0,
#                         'report': 0,
#                         'vote': 0
#                     }
#                 }
#             }

#             try:
#                 if not self.esh().es().exists(index='tracking-total', doc_type='_doc', id=self.esh().day()):
#                     self.esh().es().create(index='tracking-total', doc_type='_doc', id=self.esh().day(), body=create_tracking_element)

#                 self.esh().es().update(index='tracking-total', id=self.esh().day(), doc_type='_doc', body=script_update_registrations)
#             except Exception as err:
#                 print(err)
#                 print('Could not update tracking total')

#             if self.location():
#                 cur_lat_lng = s2.LatLng.from_degrees(self.location()['lat'], self.location()['lon'])

#                 if cur_lat_lng.is_valid():
#                     level16 = s2.CellId.from_lat_lng(cur_lat_lng).parent(16).to_token()

#                     create_tracking_element['location'] = self.location()
#                     create_tracking_element['cellId'] = level16
#                     try:
#                         if not self.esh().es().exists(index='tracking-cell', doc_type='_doc', id=self.esh().day() + '_' + level16):
#                             self.esh().es().create(index='tracking-cell', doc_type='_doc', id=self.esh().day() + '_' + level16, body=create_tracking_element)

#                         self.esh().es().update(index='tracking-cell', id=self.esh().day() + '_' + level16, doc_type='_doc', body=script_update_registrations)
#                     except Exception as err:
#                         print(err)
#                         print('Could not update tracking cell')


#                     unitid = None
#                     try:
#                         college_hits = self.esh().es().search(index='college-info', doc_type='_doc', body=TemplateCollegePointQuery(esh=self.esh(), location=self.location()).query_college_point())
#                         if college_hits.get('hits') and college_hits['hits']['total'] >= 1:
#                             unitid = college_hits['hits']['hits'][0]['_source']['unitid']
#                     except Exception as err:
#                         print(err)
#                         print('Could not search for colleges')

#                     if unitid:
#                         create_tracking_element['locationId'] = unitid
#                         try:
#                             if not self.esh().es().exists(index='tracking-location', doc_type='_doc', id=self.esh().day() + '_' + unitid):
#                                 self.esh().es().create(index='tracking-location', doc_type='_doc', id=self.esh().day() + '_' + unitid, body=create_tracking_element)

#                             self.esh().es().update(index='tracking-location', id=self.esh().day() + '_' + unitid, doc_type='_doc', body=script_update_registrations)
#                         except Exception as err:
#                             print(err)
#                             print('Could not update tracking location')

class ESHelper(object):

    def __init__(self, event=None):
        credentials = Session().get_credentials()
        es_auth = AWS4Auth(credentials.access_key, credentials.secret_key, REGION, 'es', session_token=credentials.token)
        es_headers = { "Content-Type": "application/json" }

        self.__es = Elasticsearch([ESHOST], http_auth=es_auth, use_ssl=True, verify_certs=True, connection_class=RequestsHttpConnection)

        if event and 'body' in event and isinstance(event['body'], str):
            event['body'] = json.loads(event['body'])

        self.__event = event
        self.__created_at = datetime.datetime.utcnow()
        self.__created_at_milliseconds = int((self.this_datetime() - datetime.datetime(1970, 1, 1)).total_seconds() * 1000)
        self.__created_at_hour = str(self.this_datetime().replace(microsecond=0,second=0,minute=0))
        self.__created_at_day = self.this_datetime().replace(microsecond=0,second=0,minute=0, hour=0).strftime('%Y-%m-%d')
        self.__created_at_day_before = (self.this_datetime() - datetime.timedelta(days=1)).replace(microsecond=0,second=0,minute=0, hour=0).strftime('%Y-%m-%d')
        self.__created_at_week = (self.this_datetime() - datetime.timedelta(days=(0-self.this_datetime().weekday()+7)%7)).strftime('%Y-%m-%d')

    def es(self):
        return self.__es

    def return_code(self, code=400, message=None):
        return {'statusCode': code, 'headers': {'Content-Type': 'application/json'}, 'body': json.dumps({'staging': 'production', 'version': '11-27-2018', 'statusCode': code, 'message': message})}

    def return_code_with(self, code=400, message=None, elements=None):
        rc = self.return_code(code=code, message=message)
        rc['body'] = json.loads(rc['body'])
        for key in elements:
            rc['body'][key] = elements[key]
        rc['body'] = json.dumps(rc['body'])
        return rc
    
    def this_datetime(self):
        return self.__created_at

    def milliseconds(self):
        return self.__created_at_milliseconds

    def hour(self):
        return self.__created_at_hour

    def day_before(self):
        return self.__created_at_day_before

    def day(self):
        return self.__created_at_day

    def week(self):
        return self.__created_at_week

    def event(self):
        return self.__event

    def validate_event(self):
        assert 'body' in self.event()

        body = self.event()['body']

        assert 'version' in body and 'staging' in body and 'elements' in body, 'Not implemented'

        assert body['version'] == VERSION, 'incorrect version'

        assert body['staging'] == STAGING, 'incorrect staging'

