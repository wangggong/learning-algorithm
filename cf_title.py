# -*- encoding: utf-8 -*-

import requests
import sys
import os
from bs4 import BeautifulSoup
from bs4.element import Tag


def get_question_info(ctx):
    '''get question infomation from the url'''
    base_url = 'https://codeforces.com/problemset/problem/{}/{}'
    url = base_url.format(
        ctx['base_info']['contest'],
        ctx['base_info']['problem_number'],
    )

    r = requests.get(url)
    assert r.status_code == 200
    soup = BeautifulSoup(r.text, 'lxml')

    return {
        'base_info': ctx['base_info'],
        'title': get_title(soup),
        'detail': get_detail(soup),
    }


def get_title(soup):
    '''get title info from the content'''
    title_tag = soup.find(class_='title')
    assert title_tag is not None
    title = title_tag.text
    file_title = '_'.join([x.lower() for x in title.split()[1:]])
    return {
        'title': title,
        'file_title': file_title,
    }


def get_detail(soup):
    '''get problem detail from the content'''
    selector = '#pageContent > div.problemindexholder > div.ttypography > div'
    detail_tag = soup.select(selector)
    assert len(detail_tag) > 0 and isinstance(detail_tag[0], Tag)
    # 懒得处理了, 直接上吧
    return detail_tag[0]


def parse_info(detail):
    '''parse infomation into detail strings'''
    assert detail is not None and isinstance(detail, Tag)
    return '\n'.join(_parse_info(detail, 0))


def _parse_info(detail, k):
    assert detail is not None and (
        hasattr(detail, 'text') or hasattr(detail, 'children')
    )
    if not hasattr(detail, 'children'):
        return ['  ' * k + detail.text]
    infos = []
    for child in detail.children:
        infos += _parse_info(child, k + 1)
    return infos


def gen_file(ctx, info):
    '''generate file based on the info'''
    detail = '/*\n{}\n */'.format(parse_info(info['detail']))
    with open('cf{}{}_{}.{}'.format(
        info['base_info']['contest'],
        info['base_info']['problem_number'],
        info['title']['file_title'],
        ctx['lang'],
    ), 'x') as f:
        if ctx['write_detail']:
            f.write(detail)


def main():
    assert len(sys.argv) >= 3
    contest, problem_number = sys.argv[1], sys.argv[2]
    lang = sys.argv[3] if len(sys.argv) >= 4 else 'cpp'
    ctx = {
        'base_info': {
            'contest': contest,
            'problem_number': problem_number,
        },
        'lang': lang,
        'write_detail': False,
    }
    info = get_question_info(ctx)
    gen_file(ctx, info)


if __name__ == '__main__':
    main()
