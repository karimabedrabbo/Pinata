# -*- coding: utf-8 -*-

from eshelper import ESHelper, UtilitiesTracking

def lambda_handler(event, context):

	esh = ESHelper(event)

	user_id = event['request']['userAttributes']['phone_number']

	user_hit = None
	try:
		user_hit = esh.es().get(index='users', id=user_id, doc_type='_doc')
	except Exception as err:
		print(err)
		print('Could not fetch user')

	if not user_hit or not user_hit.get('found'):
		print('User not found')
		raise ValueError('We\'re a little busy. Please try again later.')

	es_script_update_verify = {
		'script': {
			'source': 'ctx._source.properties.unverified = params.unverified',
			'lang': 'painless',
			'params': {
				'unverified': False
			}
		}
	}

	try:
		esh.es().update(index='users', doc_type='_doc', id=user_id, body=es_script_update_verify)
	except Exception as err:
		print(err)
		print('Could not update unverified attribute when signing in')

	eshtracking = UtilitiesTracking(esh=esh, user_id=user_id, user=user_hit['_source'], active=True)
		
	try:
		eshtracking.update_tracking()
	except AssertionError as err:
		print(err)

	return event

