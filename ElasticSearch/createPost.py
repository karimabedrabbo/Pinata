# -*- coding: utf-8 -*-

from eshelper import ESHelper, UtilitiesTracking, TemplatePost, S3BUCKET

from boto3 import Session

import json


def lambda_handler(event, context):
	s3 = Session().resource('s3')

	esh = ESHelper(event)

	try:
		esh.validate_event()
	except AssertionError as err:
		return esh.return_code(message=err, code=500)


	post = esh.event()['body']['elements'].get('post')


	if not post or not post.get('properties') or not post.get('properties').get('userId') or not post.get('properties').get('location'):
		return esh.return_code(message='Missing Information', code=404)

	user_id = post['properties']['userId']
	location = post['properties']['location']

	videos = post.get('video')
	images = post.get('image')
	texts = post.get('text')

	post_ids = {}
	for post_type in [videos, images, texts]:
		if post_type:
			for type_element in post_type:
				pId = type_element.get('postId')
				cId = type_element.get('childId')
				if cId:
					corr_post = post_ids.get(pId)
					children = corr_post.get('children')
					if children:
						children.append(type_element)
						corr_post['children'] = children
					else:
						corr_post['children'] = [type_element]
				elif pId:
					post_ids[pId] = type_element
					
	for pId in post_ids:
		post_element = post_ids.get(pId)

		s3object = s3.Object(S3BUCKET, 'public' + '/' + pId + '/' + 'json' + '/' + pId + '.json')

		try:
			s3object.put(
				Body=(bytes(json.dumps(post_element).encode('utf8')))
			)
		except Exception as err:
			print(err)
			print('Could not put json into s3')

		postgen = TemplatePost(esh=esh, post_id=pId, post=post_element, user_id=user_id, location=location)

		try:
			esh.es().create(index='posts', id=pId, doc_type='_doc', body=postgen.gen_create_post())
		except Exception as err:
			print(err)
			return esh.return_code(code=400, message='Could not post item')


		try:
			esh.es().index(index='events', doc_type='_doc', body=postgen.gen_create_event_location())
		except Exception as err:
			print(err)
			return esh.return_code(code=400, message='Could not post item')

		eshtracking = UtilitiesTracking(esh=esh, user_id=user_id, location=location, post=True, active=True, engagement=True)
		
		try:
			eshtracking.update_tracking()
		except AssertionError as err:
			print(err)

		return esh.return_code(code=200, message='OK!')