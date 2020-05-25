# frozen_string_literal: true

require 'spec_helper'

RSpec.describe 'parsing query' do
  it 'can parse ranges' do
    expect(p('price:2000').price).to eq(2000..)
    expect(p('price:>2000').price).to eq(2000..)

    expect(p('price:2000-3000').price).to eq(2000..3000)
    expect(p('price:-2000').price).to eq(0..2000)
    expect(p('price:<2000').price).to eq(0..2000)
  end

  it 'can have booleans' do
    expect(p('has-picture').has_picture).to be_truthy
    expect(p('has-picture:true').has_picture).to be_truthy
    expect(p('has-picture:false').has_picture).to be_falsy
  end

  it 'handles strings' do
    expect(p('q:"my name is earl"').q).to eq 'my name is earl'
    expect(p("q:'my name is earl'").q).to eq 'my name is earl'
  end

  it 'handles multiple types' do
    params = p('price:2000 has-picture q:"my name is earl"')
    expect(params.price).to eq(2000..)
    expect(params.has_picture).to be_truthy
    expect(params.q).to eq 'my name is earl'
  end
end
