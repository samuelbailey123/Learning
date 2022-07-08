from aiohttp import web
import os
import asyncio
from aiohttp.web_request import Request
import psycopg2
from flask_restful import reqparse, abort

app = web.Application()
routes = web.RouteTableDef()
AWS_PASSWORD = os.getenv('AWS_PASSWORD')
AWS_USERNAME = os.getenv('AWS_USERNAME')
AWS_PORT = os.getenv('AWS_PORT')
AWS_ENDPOINT = os.getenv('AWS_ENDPOINT')
AWS_DATABASE_TYPE = os.getenv('AWS_DATABASE_TYPE')
parser = reqparse.RequestParser()

conn = psycopg2.connect(
            host=AWS_ENDPOINT,
            database=AWS_DATABASE_TYPE,
            user=AWS_USERNAME,
            password=AWS_PASSWORD
)
cursor = conn.cursor()

def parse_record_body(args):
    return {'title': args['title'], 'description': args['description'], 'author':args['author'], 'body': args['body']}

def does_post_exist(post_id): 
    try:
        cursor.execute("SELECT * FROM posts WHERE post_id = %s;", str(post_id))
        a = cursor.fetchall()
        return True
    except (Exception, psycopg2.Error) as error:
        return abort(404, message='This post doees not exist')

@routes.view('/posts')
class posts(web.View): 
        
    async def get(self):
        post_data = []
        cursor.execute("SELECT * FROM posts;")
        raw_posts = cursor.fetchall()
        cursor.execute("SELECT count(*) AS exact_count FROM posts;")
        amount_of_posts = cursor.fetchall()[0][0]
        for row in range(0,amount_of_posts):
            post_add = {'post_id' : raw_posts[row][0], 'title' : raw_posts[row][1], 'description' : raw_posts[row][2], 'author' : raw_posts[row][3], 'body' : raw_posts[row][4], 'date' : raw_posts[row][5]}
            post_data.append(post_add)
            fill_data = post_data
        del post_data
        return web.json_response(amount_of_posts)

    async def post(self): 
        try: 
            post_input = parse_record_body(parser.parse_args())
            input_title = post_input['title']
            input_description = post_input['description']
            input_author = post_input['author']
            input_body = post_input['body']
            cursor.execute("INSERT INTO posts (title, description, author, body) VALUES(%s, %s, %s, %s);", (input_title, input_description, input_author, input_body))
            conn.commit()
            return "Succesfully added into table"
        except (Exception, psycopg2.Error) as error:
            return "Failed inserting record into table {}".format(error)
    
@asyncio.coroutine
def post_get(request):
    post_id = Request.match_info['number']
    if does_post_exist(post_id): 
            cursor.execute("SELECT * FROM posts WHERE post_id = %s;", (str(post_id)))
            return_data = cursor.fetchall()
            if not return_data: 
                abort(404, message='This post doees not exist')
            else: 
                return web.json_response(return_data)

def delete(self):
    post_id = Request.match_info['number']
    does_post_exist(post_id)
    cursor.execute("DELETE FROM posts WHERE post_id = %s;", (str(post_id)))
    return ("The post has been deleted")

def put(self): 
    post_id = Request.match_info['number']
    post_input = parse_record_body(parser.parse_args())
    input_title = post_input['title']
    input_description = post_input['description']
    input_author = post_input['author']
    input_body = post_input['body']
    try:
        if input_title: 
            cursor.execute("UPDATE posts SET title = %s WHERE post_id = %s;", (input_title, post_id))
        if input_description: 
            cursor.execute("UPDATE posts SET description = %s WHERE post_id = %s;", (input_description, post_id))
        if input_author: 
            cursor.execute("UPDATE posts SET author = %s WHERE post_id = %s;", (input_author, post_id))
        if input_body: 
            cursor.execute("UPDATE posts SET body = %s WHERE post_id = %s;", (input_body, post_id))
        return ('Succesful update')
    except (Exception, psycopg2.Error) as error:
        return "Failed inserting into table {}".format(error)

app.router.add_route('GET', '/posts/{number:\d+}', post_get)
app.router.add_routes(routes)

web.run_app(app)
